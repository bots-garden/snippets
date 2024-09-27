#!/bin/bash
go run snippets.go generate --input samples/js.yml --output ../.vscode/js.code-snippets 
go run snippets.go generate --input samples/js.yml --output samples/js.code-snippets 

