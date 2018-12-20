FROM golang:1.11

LABEL version="1.0.0"
LABEL maintainer="Apex"
LABEL repository="http://github.com/apex/actions"
LABEL homepage="http://github.com/apex/actions/up"
LABEL "com.github.actions.name"="Up"
LABEL "com.github.actions.description"="Perform Up application operations"
LABEL "com.github.actions.icon"="chevron-up"
LABEL "com.github.actions.color"="white"

ENV CI true
RUN curl -sf https://up.apex.sh/install | sh
RUN chmod +x /usr/local/bin/up

RUN go get github.com/apex/actions/up/cmd/up-wrapper

ENTRYPOINT ["up-wrapper"]
CMD ["deploy", "--no-build"]
