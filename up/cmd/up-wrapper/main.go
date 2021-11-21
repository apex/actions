package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/apex/actions/slack"
)

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
			IconURL:  "https://avatars.slack-edge.com/2018-12-20/508671226196_a96b52b97348bd9675e2_192.png",
			Attachments: []*slack.Attachment{
				&slack.Attachment{
					Title:  os.Getenv("GITHUB_REPOSITORY"),
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