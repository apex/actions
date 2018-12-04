FROM alpine:3.8

LABEL version="1.0.0"
LABEL maintainer="Apex"
LABEL repository="http://github.com/apex/actions"
LABEL homepage="http://github.com/apex/actions/up"
LABEL "com.github.actions.name"="Up Deploy"
LABEL "com.github.actions.description"="Deploy an Up application"
LABEL "com.github.actions.icon"="upload-cloud"
LABEL "com.github.actions.color"="purple"

RUN apk add --update curl jq && rm -rf /var/cache/apk/*

RUN curl -sf https://up.apex.sh/install | sh
RUN chmod +x /usr/local/bin/up

ENTRYPOINT ["/usr/local/bin/up"]
