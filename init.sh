#!/bin/bash

chmod +x .githooks/commit-msg
chmod +x .githooks/pre-push
git config core.hooksPath .githooks
cd api
go mod init github.com/Paulehair/elasticsearch-golang-project
go get