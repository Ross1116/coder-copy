# Copy Comment Remover

A simple utility that automatically removes comments from code in your clipboard.

## Overview

This tool monitors your clipboard for code snippets and automatically removes:
- Single-line comments (`//`)
- Multi-line comments (`/* */`)

The cleaned code is then placed back into your clipboard, ready to paste without comments.

## Features

- Lightweight clipboard monitoring
- Automatic comment removal
- Supports both single-line and multi-line comment styles
- Zero configuration required

## Usage

The application is already built and available in the `bin` directory:

```bash
./bin/copy-comment-remover
```

After starting the application:
1. Copy any code containing comments
2. The application automatically processes the code
3. Paste the cleaned code wherever you need it

## Building and Running

If you need to rebuild the application, this project requires CGO to be enabled for clipboard functionality.

### Using Make

A Makefile is provided for convenience:

```bash
# Build the application
make build

# Run the application
make run

# Clean build artifacts
make clean
```

### Manual Build

If you prefer not to use Make:

```bash
# Enable CGO (required for clipboard functionality)
export CGO_ENABLED=1

# Build
go build -o bin/copy-comment-remover cmd/app/main.go

# Run
go run cmd/app/main.go
```

## How It Works

The tool uses regular expressions to identify and remove comments from code. It watches the clipboard for changes and only processes new content, avoiding unnecessary operations.

## Requirements

- Go 1.16+
- [golang.design/x/clipboard](https://pkg.go.dev/golang.design/x/clipboard) package
