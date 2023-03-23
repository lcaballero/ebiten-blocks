#!/bin/bash

build() {
    go generate ./... && go install
}

tests() {
    go test ./...
}

see::cover() {
  open "http://localhost:2222/cover.html#file$1"
}

cover() {
    mkdir -p .cover && \
        go test -coverprofile=.cover/cover.out \
           ./ && \
        go tool cover -html=.cover/cover.out -o .cover/cover.html
}


"$@"
