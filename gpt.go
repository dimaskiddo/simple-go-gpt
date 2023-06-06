package main

import (
	"context"
	"fmt"
	"os"

	OpenAI "github.com/sashabaranov/go-openai"
)

var OAIClient *OpenAI.Client

func init() {
	OAIAPIKey := os.Getenv("OPENAI_API_KEY")
	OAIClient = OpenAI.NewClient(OAIAPIKey)
}

func GPT3Completion(question string) (err error) {
	gptRequest := OpenAI.ChatCompletionRequest{
		Model:            "gpt-3.5-turbo",
		MaxTokens:        4000,
		Temperature:      0.0,
		TopP:             1.0,
		PresencePenalty:  0.0,
		FrequencyPenalty: 0.6,
		Messages: []OpenAI.ChatCompletionMessage{
			{
				Role:    OpenAI.ChatMessageRoleUser,
				Content: question,
			},
		},
	}

	gptResponse, err := OAIClient.CreateChatCompletion(
		context.Background(),
		gptRequest,
	)

	if err != nil {
		return err
	}

	if len(gptResponse.Choices) > 0 {
		fmt.Println(gptResponse.Choices[0].Message.Content)
	}

	return nil
}
