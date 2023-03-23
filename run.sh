#!/bin/bash

build() {
    go generate ./... && go install
}

tests() {
    go test ./...
}

"$@"
