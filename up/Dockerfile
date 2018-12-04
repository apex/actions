FROM alpine:3.8

LABEL version="1.0.0"
LABEL maintainer="Apex"
LABEL repository="http://github.com/apex/actions"
LABEL homepage="http://github.com/apex/actions/up"
LABEL "com.github.actions.name"="Up"
LABEL "com.github.actions.description"="Perform Up application operations"
LABEL "com.github.actions.icon"="arrow-up"
LABEL "com.github.actions.color"="chevron-up"

RUN apk add --update curl jq && rm -rf /var/cache/apk/*

ENV CI true
RUN curl -sf https://up.apex.sh/install | sh
RUN chmod +x /usr/local/bin/up

ENTRYPOINT ["/usr/local/bin/up"]
