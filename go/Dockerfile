FROM golang:1.11

LABEL version="1.0.0"
LABEL maintainer="Apex"
LABEL repository="http://github.com/apex/actions"
LABEL homepage="http://github.com/apex/actions/go"
LABEL "com.github.actions.name"="Go"
LABEL "com.github.actions.description"="Golang"
LABEL "com.github.actions.icon"="code"
LABEL "com.github.actions.color"="white"

ENV GO111MODULE on

ENTRYPOINT ["/usr/local/go/bin/go"]
CMD ["build", "-o", "server", "main.go"]
