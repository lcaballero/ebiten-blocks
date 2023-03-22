#!/bin/bash

build() {
    go install
}

tests() {
    go test ./...
}

"$@"
