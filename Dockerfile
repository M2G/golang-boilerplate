FROM        golang:1.26.2-alpine3.23 as base
ENV         GO111MODULE     on

RUN         apk -u add git openssh build-base
WORKDIR     /go/src/github.com/golang-boilerplate
ADD         .   /go/src/github.com/golang-boilerplate
COPY        scripts/dep .
RUN         chmod +x dep; ./dep

FROM        base as dev
EXPOSE      80 443 43554
ADD         .  /go/src/github.com/golang-boilerplate
WORKDIR     /go/src/github.com/golang-boilerplate
COPY        modd /usr/local/bin/modd
RUN         chmod +x /usr/local/bin/modd
ENTRYPOINT  ["modd"]
CMD         ["-f", "configuration/modd/modd.conf"]

FROM        base as builder
ADD         .   /go/src/github.com/golang-boilerplate
WORKDIR     /go/src/github.com/golang-boilerplate
RUN         chmod +x scripts/dep; ./scripts/dep
RUN         make build

FROM        alpine:latest as release
EXPOSE      80 443 43554
RUN         apk -u add ca-certificates
COPY        --from=builder /go/src/github.com/golang-boilerplate/bin/golang-boilerplate golang-boilerplate
ENTRYPOINT  [ "/golang-boilerplate" ]
