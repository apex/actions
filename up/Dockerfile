FROM golang:1.15

RUN curl -sL https://deb.nodesource.com/setup_11.x | bash -
RUN apt-get install -y nodejs

ENV CI true
RUN curl -sf https://up.apex.sh/install | sh
RUN chmod +x /usr/local/bin/up

RUN go get github.com/apex/actions/up/cmd/up-wrapper

ENTRYPOINT ["up-wrapper", "deploy", "--no-build"]