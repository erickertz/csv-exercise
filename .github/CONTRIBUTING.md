# Format & Linting

This repo uses goimports and gofmt for formatting. For linting GolangCI-Lint (https://github.com/golangci/golangci-lint). A GolangCI-Lint configuration file is located at the project root/golangci.yaml .

For goimports and gofmt for formatting, run the following from project root:

`goimports -e -w -l . && gofmt -e -s -w -l .`

For GolangCI-Lint linting, run the following from project root:

`golangci-lint run`
