# Coder Copy

Ever got annoyed copying code only to find it has tons of comments that you need to manually remove? This tool automatically does that for you!

## Overview

This tool monitors your clipboard for code snippets and automatically removes comments, leaving clean, comment-free code ready to paste.

## Features

- Lightweight clipboard monitoring
- Automatic comment removal
- Optional code formatting (Go only)
- Interactive TUI with easy configuration
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
./bin/coder-copy

# With command-line flags
./bin/coder-copy -go
./bin/coder-copy -python
./bin/coder-copy -js
./bin/coder-copy -jsx
./bin/coder-copy -java
./bin/coder-copy -c

# Enable auto-formatting (Go only)
./bin/coder-copy -format
```

### Interactive Mode

Without command-line arguments, the application starts in interactive mode:

1. Select your programming language using arrow keys (↑/↓) and press Enter
2. Choose whether to enable auto-formatting (currently Go only)
3. Start copying code with comments

### Navigation

- Press `s` to change settings while monitoring
- Press `backspace` to go back to previous screens
- Press `q` to quit the application

## Workflow

1. Start the application
2. Copy any code containing comments
3. The application automatically processes the code
4. Paste the cleaned code wherever you need it

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
go build -o bin/coder-copy cmd/app/main.go

# Run
go run cmd/app/main.go
```

## How It Works

The tool uses regular expressions to identify and remove comments from code. It watches the clipboard for changes and only processes new content, avoiding unnecessary operations. The interactive TUI is built using [Bubble Tea](https://github.com/charmbracelet/bubbletea).

## Requirements

- Go 1.16+
- [golang.design/x/clipboard](https://pkg.go.dev/golang.design/x/clipboard) package
- [github.com/charmbracelet/bubbletea](https://github.com/charmbracelet/bubbletea) package