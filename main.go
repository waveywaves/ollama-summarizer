package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/waveywaves/summarizer/pkg/summarizer"
)

func main() {
	// Create a new summarizer instance
	sum, err := summarizer.New(summarizer.Config{
		Model: "mistral", // Using mistral model
	})
	if err != nil {
		log.Fatalf("Failed to create summarizer: %v", err)
	}

	// Example changes to summarize
	changes := `
	- Added new authentication middleware
	- Updated user profile API endpoints
	- Fixed bug in password reset flow
	- Improved error handling in login process
	`

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Generate the summary
	summary, err := sum.SummarizeWithCustomPrompt(ctx, changes, `Summarize the following changes in a release note: %s`)
	if err != nil {
		log.Fatalf("Failed to generate summary: %v", err)
	}

	// Print the summary
	fmt.Printf("\nRelease Notes Summary:\n%s\n", summary)
}
