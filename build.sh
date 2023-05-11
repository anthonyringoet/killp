#!/bin/bash

# Build for Windows
echo "Building killp for Windows"
GOOS=windows GOARCH=amd64 go build -o killp_windows_amd64 main.go process.go

# Build for macOS
echo "Building killp for macOS"
GOOS=darwin GOARCH=amd64 go build -o killp_darwin_amd64 main.go process.go

# Build for Linux
echo "Building killp for Linux"
GOOS=linux GOARCH=amd64 go build -o killp_linux_amd64 main.go process.go
