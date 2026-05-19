FROM golang:1.26.2 AS base
ENV GO111MODULE=on

ARG MODD_VERSION=v1.0.2
ENV MODD_VERSION=${MODD_VERSION}

RUN apt-get update && apt-get install -y --no-install-recommends \
    git \
    openssh-client \
    wget \
    && rm -rf /var/lib/apt/lists/*

# Install modd from GitHub release (detect architecture)
RUN ARCH=$(uname -m) && \
    if [ "$ARCH" = "x86_64" ]; then ARCH="amd64"; elif [ "$ARCH" = "aarch64" ]; then ARCH="arm64"; fi && \
    wget -qO- https://github.com/M2G/modd/releases/download/${MODD_VERSION}/modd-${MODD_VERSION}-linux-${ARCH}.tgz | \
    tar -xz -C /usr/local/bin/ && \
    chmod +x /usr/local/bin/modd

WORKDIR /go/src/github.com/golang-boilerplate
ADD . /go/src/github.com/golang-boilerplate
COPY scripts/dep .
RUN chmod +x dep; ./dep

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

FROM debian:bookworm-slim AS release
EXPOSE 80 443 43554
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates && rm -rf /var/lib/apt/lists/*
COPY --from=builder /go/src/github.com/golang-boilerplate/bin/golang-boilerplate golang-boilerplate
ENTRYPOINT ["/golang-boilerplate"]
