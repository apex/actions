FROM golang:1.15

ENV CI true
RUN curl -sf https://up.apex.sh/install | sh
RUN chmod +x /usr/local/bin/up

RUN go get github.com/apex/actions/up/cmd/up-wrapper

ENTRYPOINT ["up-wrapper"]
CMD ["deploy", "--no-build"]
