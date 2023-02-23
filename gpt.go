package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	gpt "github.com/sashabaranov/go-gpt3"
)

var OpenAI *gpt.Client

func init() {
	gptAPIKey := os.Getenv("OPENAI_API_KEY")
	OpenAI = gpt.NewClient(gptAPIKey)
}

func GPTResponse(question string) (err error) {
	gptRequest := gpt.CompletionRequest{
		Model:            "text-davinci-003",
		MaxTokens:        4000,
		Temperature:      0,
		TopP:             1,
		PresencePenalty:  0,
		FrequencyPenalty: 0.6,
		Prompt:           question,
	}

	gptResponse, err := OpenAI.CreateCompletionStream(
		context.Background(),
		gptRequest,
	)

	if err != nil {
		return err
	}
	defer gptResponse.Close()

	wordPosition := 0
	for {
		gptResponseStream, err := gptResponse.Recv()
		if errors.Is(err, io.EOF) {
			break
		}

		if len(gptResponseStream.Choices) > 0 {
			if wordPosition == 0 {
				fmt.Printf("%v", strings.TrimLeft(gptResponseStream.Choices[0].Text, "\n"))
			} else {
				fmt.Printf("%v", gptResponseStream.Choices[0].Text)
			}

			wordPosition++
		}
	}
	fmt.Println("")

	return nil
}
