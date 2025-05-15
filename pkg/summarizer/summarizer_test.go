package summarizer

import (
	"context"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		config  Config
		wantErr bool
	}{
		{
			name: "default config",
			config: Config{
				Model: "",
			},
			wantErr: false,
		},
		{
			name: "custom model",
			config: Config{
				Model: "codellama",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil && !tt.wantErr {
				t.Error("New() returned nil but no error was expected")
			}
		})
	}
}

func TestSummarizer_SummarizeChanges(t *testing.T) {
	s, err := New(Config{})
	if err != nil {
		t.Fatalf("Failed to create summarizer: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	changes := `
	- Added new authentication middleware
	- Updated user profile API endpoints
	- Fixed bug in password reset flow
	`

	summary, err := s.SummarizeChanges(ctx, changes)
	if err != nil {
		t.Errorf("SummarizeChanges() error = %v", err)
		return
	}

	if summary == "" {
		t.Error("SummarizeChanges() returned empty summary")
	}
}

func TestSummarizer_SummarizeWithCustomPrompt(t *testing.T) {
	s, err := New(Config{})
	if err != nil {
		t.Fatalf("Failed to create summarizer: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	content := "Test content"
	promptTemplate := "Summarize this: %s"

	summary, err := s.SummarizeWithCustomPrompt(ctx, content, promptTemplate)
	if err != nil {
		t.Errorf("SummarizeWithCustomPrompt() error = %v", err)
		return
	}

	if summary == "" {
		t.Error("SummarizeWithCustomPrompt() returned empty summary")
	}
}
