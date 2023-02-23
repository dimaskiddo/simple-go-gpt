package main

import (
	"context"
	"os"
	"strings"

	gpt "github.com/sashabaranov/go-gpt3"
)

var OpenAI *gpt.Client

func init() {
	gptAPIKey := os.Getenv("OPENAI_API_KEY")
	OpenAI = gpt.NewClient(gptAPIKey)
}

func GPTResponse(question string) (response string, err error) {
	gptRequest := gpt.CompletionRequest{
		Model:            gpt.GPT3TextDavinci003,
		MaxTokens:        4000,
		Temperature:      0,
		TopP:             1,
		PresencePenalty:  0,
		FrequencyPenalty: 0,
		Prompt:           question,
	}

	gptResponse, err := OpenAI.CreateCompletion(
		context.Background(),
		gptRequest,
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
