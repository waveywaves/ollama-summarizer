# Summarizer

A Go package that generates release notes and pull request summaries using Ollama-hosted AI models.

## Installation

```bash
go get github.com/waveywaves/summarizer
```

## Prerequisites

1. Install [Go](https://golang.org/doc/install) (version 1.24 or later)
2. Install [Ollama](https://ollama.ai/)
3. Pull the mistral model:
   ```bash
   ollama pull mistral
   ```

## Usage

### As a Package

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/waveywaves/summarizer/pkg/summarizer"
)

func main() {
    // Create a new summarizer
    sum, err := summarizer.New(summarizer.Config{
        Model: "mistral", // Optional: defaults to "mistral"
    })
    if err != nil {
        log.Fatal(err)
    }

    // Create context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    // Example changes
    changes := `
    - Added user authentication
    - Fixed bug in login flow
    `

    // Generate summary
    summary, err := sum.SummarizeChanges(ctx, changes)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(summary)
}
```

### Using Custom Prompts

You can also use custom prompt templates:

```go
summary, err := sum.SummarizeWithCustomPrompt(ctx, content, "Please analyze this code: %s")
```

### Running the Example

1. Start Ollama in the background:
   ```bash
   ollama serve
   ```

2. Run the example program:
   ```bash
   go run main.go
   ```

## Configuration

The package can be configured using the `Config` struct:

```go
type Config struct {
    Model     string // The Ollama model to use (default: "mistral")
    OllamaURL string // Optional: URL for the Ollama API
}
```

## Why Mistral?

The package defaults to using the Mistral model because:
- It excels at understanding code context and technical content
- It can process long sequences of changes efficiently
- It maintains good context awareness across multiple files
- It produces well-structured, coherent summaries

## Alternative Models

While the default is Mistral, you can use other models:
- codellama (good for code-specific tasks)
- llama2 (general purpose)
- neural-chat (conversational)

## Running Tests

```bash
go test ./pkg/summarizer -v
```

## License

This project is licensed under the MIT License - see the LICENSE file for details. # ollama-summarizer
