FROM mcr.microsoft.com/devcontainers/go:1.27

RUN go install golang.org/x/tools/gopls@latest && \
  go install github.com/go-delve/delve/cmd/dlv@latest && \
  go install github.com/air-verse/air@latest && \
  go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

WORKDIR /workspace

USER vscode

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN mkdir -p /home/vscode/.air && \
  echo '[build]\n  cmd = "go build -o ./tmp/main ./cmd/main.go"\n  bin = "./tmp/main"\n  include_ext = ["go"]\n  exclude_dir = ["tmp", "vendor"]\n[misc]\n  clean_on_exit = true' > /home/vscode/.air/air.toml
