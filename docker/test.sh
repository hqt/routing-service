#!/usr/bin/env bash

golangci-lint run ./... -v  \
&& go test ./... -race -p 1 -coverprofile=cover.out \
&& go clean -testcache ./... \
&& go tool cover -func output/cover.out | grep total | awk '{print $3}' > output/coverage.txt
