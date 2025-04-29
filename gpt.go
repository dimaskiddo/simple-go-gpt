package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	OpenAI "github.com/sashabaranov/go-openai"
)

var OAIClient *OpenAI.Client

func init() {
	OAIAPIKey := os.Getenv("OPENAI_API_KEY")

	OAIConfig := OpenAI.DefaultConfig(OAIAPIKey)
	OAIConfig.BaseURL = os.Getenv("OPENAI_API_URL")

	OAIClient = OpenAI.NewClientWithConfig(OAIConfig)
}

func GPT3Completion(question string) (err error) {
	isStream := new(bool)
	*isStream = true

	var OAIGPTResponseText string

	OAIGPTChatCompletion := []OpenAI.ChatCompletionMessage{
		{
			Role:    OpenAI.ChatMessageRoleUser,
			Content: question,
		},
	}

	OAIGPTPrompt := OpenAI.ChatCompletionRequest{
		Model:            os.Getenv("OPENAI_MODEL_NAME"),
		MaxTokens:        4096,
		Temperature:      0.8,
		TopP:             0.9,
		PresencePenalty:  0.0,
		FrequencyPenalty: 0.0,
		Messages:         OAIGPTChatCompletion,
		Stream:           *isStream,
	}

	OAIGPTStream, err := OAIClient.CreateChatCompletionStream(
		context.Background(),
		OAIGPTPrompt,
	)

	if err != nil {
		return err
	}
	defer OAIGPTStream.Close()

	for {
		OAIGPTResponse, err := OAIGPTStream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return err
		}

		if len(OAIGPTResponse.Choices) > 0 {
			OAIGPTResponseText = OAIGPTResponseText + OAIGPTResponse.Choices[0].Delta.Content
		}
	}

	ThinkingResponseRegex := regexp.MustCompile("[\\s\\S]*<\\/think>\\n?")
	CleanThinkingResponse := ThinkingResponseRegex.ReplaceAllString(OAIGPTResponseText, "")

	OAIGPTResponseBuffer := strings.TrimSpace(CleanThinkingResponse)
	OAIGPTResponseBuffer = strings.TrimLeft(OAIGPTResponseBuffer, "?\n")
	OAIGPTResponseBuffer = strings.TrimLeft(OAIGPTResponseBuffer, "!\n")
	OAIGPTResponseBuffer = strings.TrimLeft(OAIGPTResponseBuffer, ":\n")
	OAIGPTResponseBuffer = strings.TrimLeft(OAIGPTResponseBuffer, "'\n")
	OAIGPTResponseBuffer = strings.TrimLeft(OAIGPTResponseBuffer, ".\n")
	OAIGPTResponseBuffer = strings.TrimLeft(OAIGPTResponseBuffer, "\n")

	fmt.Println(OAIGPTResponseBuffer)

	return nil
}
