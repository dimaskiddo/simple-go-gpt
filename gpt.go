package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	OpenAI "github.com/sashabaranov/go-openai"
)

var GPT *OpenAI.Client

func init() {
	gptAPIKey := os.Getenv("OPENAI_API_KEY")
	GPT = OpenAI.NewClient(gptAPIKey)
}

func GPT3Completion(question string) (err error) {
	gptRequest := OpenAI.ChatCompletionRequest{
		Model:            "gpt-3.5-turbo",
		MaxTokens:        4000,
		Temperature:      0,
		TopP:             1,
		PresencePenalty:  0,
		FrequencyPenalty: 0.6,
		Messages: []OpenAI.ChatCompletionMessage{
			{
				Role:    OpenAI.ChatMessageRoleUser,
				Content: question,
			},
		},
	}

	gptResponse, err := GPT.CreateChatCompletionStream(
		context.Background(),
		gptRequest,
	)

	if err != nil {
		return err
	}
	defer gptResponse.Close()

	isFirstWordFound := false
	for {
		gptResponseStream, err := gptResponse.Recv()
		if errors.Is(err, io.EOF) {
			break
		}

		if len(gptResponseStream.Choices) > 0 {
			wordResponse := gptResponseStream.Choices[0].Delta.Content
			if !isFirstWordFound && wordResponse != "\n" && len(strings.TrimSpace(wordResponse)) != 0 {
				isFirstWordFound = true
			}

			if isFirstWordFound {
				fmt.Printf("%v", wordResponse)
			}
		}
	}

	fmt.Println("")

	return nil
}
