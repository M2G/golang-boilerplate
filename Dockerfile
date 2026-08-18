FROM golang:1.25-alpine3.22 AS fswatch

RUN apk add --no-cache git autoconf automake libtool gettext gettext-dev make g++ texinfo

WORKDIR /tmp
RUN git clone --depth 1 --branch 1.18.3 https://github.com/emcrisostomo/fswatch.git

WORKDIR /tmp/fswatch
RUN ./autogen.sh && ./configure --disable-nls && make -j && make install


FROM golang:1.25-alpine3.22 AS base

ENV CGO_ENABLED=1
ENV LD_LIBRARY_PATH=/usr/local/lib
ENV GOTOOLCHAIN=auto

RUN apk -u add git openssh build-base

COPY --from=fswatch /usr/local/lib/ /usr/local/lib/
COPY --from=fswatch /usr/local/include/ /usr/local/include/

WORKDIR /go/src/github.com/golang-boilerplate
ADD . /go/src/github.com/golang-boilerplate


FROM base AS dev

EXPOSE 80 443 43554
ADD . /go/src/github.com/golang-boilerplate
WORKDIR /go/src/github.com/golang-boilerplate
RUN go install github.com/M2G/modd/cmd/modd@v0.1.2
RUN mv "$(go env GOPATH)/bin/modd" /usr/local/bin/modd
ENTRYPOINT ["modd"]
CMD ["-f", "configuration/modd/modd.conf"]


FROM base AS builder

ADD . /go/src/github.com/golang-boilerplate
WORKDIR /go/src/github.com/golang-boilerplate
RUN make build


FROM alpine:latest AS release

EXPOSE 80 443 43554
ENV LD_LIBRARY_PATH=/usr/local/lib
RUN apk -u add ca-certificates
COPY --from=fswatch /usr/local/lib/ /usr/local/lib/
COPY --from=builder /go/src/github.com/golang-boilerplate/bin/golang-boilerplate golang-boilerplate
ENTRYPOINT [ "/golang-boilerplate" ]