package main

import (
	"context"
	"os"
	"strings"

	gpt "github.com/PullRequestInc/go-gpt3"
)

var OpenAI gpt.Client

func init() {
	gptAPIKey := os.Getenv("OPENAI_API_KEY")
	OpenAI = gpt.NewClient(gptAPIKey)
}

func ChatGPTResponse(question string) (response string, err error) {
	gptResponse, err := OpenAI.CompletionWithEngine(
		context.Background(),
		"text-davinci-003",
		gpt.CompletionRequest{
			Prompt:      []string{question},
			MaxTokens:   gpt.IntPtr(4000),
			Temperature: gpt.Float32Ptr(1.0),
			TopP:        gpt.Float32Ptr(0.9),
		},
	)

	if err != nil {
		return "", err
	}

	buffResponse := strings.TrimSpace(gptResponse.Choices[0].Text)
	buffResponse = strings.TrimLeft(buffResponse, "?\n")
	buffResponse = strings.TrimLeft(buffResponse, "!\n")
	buffResponse = strings.TrimLeft(buffResponse, ":\n")
	buffResponse = strings.TrimLeft(buffResponse, "'\n")
	buffResponse = strings.TrimLeft(buffResponse, ".\n")
	buffResponse = strings.TrimLeft(buffResponse, "\n")

	return buffResponse, nil
}
