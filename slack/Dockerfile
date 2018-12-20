FROM golang:1.11

LABEL version="1.0.0"
LABEL maintainer="Apex"
LABEL repository="http://github.com/apex/actions"
LABEL homepage="http://github.com/apex/actions/slack"
LABEL "com.github.actions.name"="Slack"
LABEL "com.github.actions.description"="Send a Slack message"
LABEL "com.github.actions.icon"="message-circle"
LABEL "com.github.actions.color"="white"

RUN go get github.com/apex/actions/slack/cmd/slack

ENTRYPOINT ["slack"]
