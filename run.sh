#!/bin/bash
go run main.go generate --input samples/js.yml --output ../.vscode/js.code-snippets 
go run main.go generate --input samples/js.yml --output samples/js.code-snippets 

