package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/apex/actions/slack"
)

// Config is the Up configuration.
type Config struct {
	Name string `json:"name"`
}

func main() {
	args := os.Args[1:]

	workdir, err := os.Getwd()
	if err != nil {
		log.Fatalf("error getting working directory: %s", err)
	}

	// chdir
	for i, arg := range args {
		if arg == "-C" {
			os.Chdir(args[i+1])
			args = args[i+2:]
		}
	}

	// is it a deployment?
	deploy := len(args) > 0 && args[0] == "deploy"

	// determine the stage
	stage := "staging"
	if deploy {
		for _, arg := range args[1:] {
			if !strings.HasPrefix(arg, "--") {
				stage = arg
				break
			}
		}
	}

	// read app name
	f, err := os.Open("up.json")
	if err != nil {
		log.Fatalf("error opening up.json: %s", err)
	}
	defer f.Close()

	var c Config
	err = json.NewDecoder(f).Decode(&c)
	if err != nil {
		log.Fatalf("error reading up.json: %s", err)
	}

	// proxy to up(1)
	start := time.Now()

	cmd := exec.Command("up", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	if deploy {
		err := slack.WriteMessage(filepath.Join(workdir, "slack.json"), &slack.Message{
			Username: "Up",
			Attachments: []*slack.Attachment{
				&slack.Attachment{
					Title:  c.Name,
					Text:   fmt.Sprintf("Deployment to *%s* completed.", stage),
					Footer: fmt.Sprintf("Completed in %s", time.Since(start).Round(time.Second)),
				},
			},
		})

		if err != nil {
			log.Fatalf("error writing slack message: %s", err)
		}
	}
}
