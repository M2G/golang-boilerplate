FROM golang:1.26.2-alpine3.23 as base
ENV GO111MODULE=on

RUN apk add --no-cache \
    git \
    openssh \
    build-base \
    autoconf \
    automake \
    libtool \
    curl \
    bash \
    linux-headers \
    gettext-dev

# Build libfswatch from source
RUN git clone --depth=1 https://github.com/emcrisostomo/fswatch /tmp/fswatch && \
    cd /tmp/fswatch && \
    ./autogen.sh && \
    ./configure --prefix=/usr/local --enable-static --disable-shared && \
    make -j$(nproc) && \
    make install

# Verify headers are in place
RUN ls /usr/local/include/libfswatch/c/libfswatch.h

WORKDIR /go/src/github.com/golang-boilerplate
ADD . /go/src/github.com/golang-boilerplate
COPY scripts/dep .
RUN chmod +x dep; ./dep

# Build modd from source with libfswatch
RUN cd tools/modd && \
    CGO_ENABLED=1 \
    CGO_CFLAGS="-I/usr/local/include" \
    CGO_LDFLAGS="-L/usr/local/lib -lfswatch -lstdc++ -lintl" \
    go build -o /usr/local/bin/modd ./cmd/modd

FROM base AS dev
EXPOSE 80 443 43554
ADD . /go/src/github.com/golang-boilerplate
WORKDIR /go/src/github.com/golang-boilerplate
ENTRYPOINT ["modd"]
CMD ["-f", "configuration/modd/modd.conf"]

FROM base AS builder
ADD . /go/src/github.com/golang-boilerplate
WORKDIR /go/src/github.com/golang-boilerplate
RUN chmod +x scripts/dep; ./scripts/dep
RUN make build

FROM alpine:latest AS release
EXPOSE 80 443 43554
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/golang-boilerplate/bin/golang-boilerplate golang-boilerplate
ENTRYPOINT ["/golang-boilerplate"]
