#!/bin/bash
go run snippets.go generate \
  --url https://raw.githubusercontent.com/bots-garden/golab-demos/main/snippets-demo/messages.yml \
  --input samples/messages.yml \
  --output ../.vscode/messages.code-snippets 

