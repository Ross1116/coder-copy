# Copy Comment Remover

Ever got annoyed copying code only to find it has tons of comments that you need to manually remove? This tool automatically does that for you!

## Overview

This tool monitors your clipboard for code snippets and automatically removes comments, leaving clean, comment-free code ready to paste.

## Features

- Lightweight clipboard monitoring
- Automatic comment removal
- Supports multiple programming languages:
  - Go (default)
  - C/C++
  - Java
  - JavaScript/TypeScript
  - JSX/TSX/React
  - Python
- Zero configuration required

## Usage

The application is already built and available in the `bin` directory:

```bash
# Default (Go comments)
./bin/copy-comment-remover
#           or 
./bin/copy-comment-remover -go

# For Python comments
./bin/copy-comment-remover -python

# For JavaScript comments
./bin/copy-comment-remover -js

# For JSX/React comments
./bin/copy-comment-remover -jsx

# For Java comments
./bin/copy-comment-remover -java

# For C/C++ comments
./bin/copy-comment-remover -c
```

After starting the application:
1. Copy any code containing comments
2. The application automatically processes the code
3. Paste the cleaned code wherever you need it

## Supported Comment Types

- Go/C/Java/JavaScript/React:
  - Single-line comments (`//`)
  - Multi-line comments (`/* */`)
  - JSX comments (`{/* */}`) in React code
- Python:
  - Single-line comments (`#`)
  - Triple-quoted docstrings (`'''` and `"""`)

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
