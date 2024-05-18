#!/bin/bash

# Project name
PROJECT_NAME="cloneall"

# List of OS/Arch combinations to build for
PLATFORMS=("windows/amd64" "windows/386" "linux/amd64" "linux/386" "darwin/amd64" "darwin/arm64")

# Build the project for each platform
for PLATFORM in "${PLATFORMS[@]}"; do
    OS=$(echo $PLATFORM | cut -d '/' -f 1)
    ARCH=$(echo $PLATFORM | cut -d '/' -f 2)
    OUTPUT_NAME=$PROJECT_NAME'-'$OS'-'$ARCH

    if [ $OS = "windows" ]; then
        OUTPUT_NAME+='.exe'
    fi

    echo "Building for $OS/$ARCH..."
    env GOOS=$OS GOARCH=$ARCH go build -o $OUTPUT_NAME cloneall.go
done

echo "Builds completed!"
