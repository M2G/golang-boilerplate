# GOLANG INIT

Sample starter operation using Golang.

## Prerequisites

Docker, Git, Go. Check `conf-dev.yml` from a configuration example.

### Installing

clone the repo, then fetch dependencies and u're good to go.

```
    git clone git@github.com/golang-starter.git $GOPATH/src/github.com/golang-starter
    cd $GOPATH/src/github.com/golang-starter
    export GO111MODULE=on;
    go mod vendor;
    go mod download;
    go mod tidy
    make dev
```

### Build locally

```
    git clone git@github.com/golang-starter.git $GOPATH/src/github.com/golang-starter
    cd $GOPATH/src/github.com/golang-starter
    export GO111MODULE=on;
    go mod vendor;
    go mod download;
    go mod tidy
    make build
```


### Running the tests

```
    export GO111MODULE=on;
    go mod vendor;
    go mod download;
    go mod tidy
    make test
```