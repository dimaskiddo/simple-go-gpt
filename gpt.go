package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"

	Ollama "github.com/ollama/ollama/api"
	OpenAI "github.com/sashabaranov/go-openai"
)

var OAIClient *OpenAI.Client
var OClient *Ollama.Client

var GPTEngine string

func init() {
	GPTEngine = os.Getenv("GPT_ENGINE")

	switch strings.ToLower(GPTEngine) {
	case "openai":
		OAIAPIKey := os.Getenv("OPENAI_API_KEY")
		OAIClient = OpenAI.NewClient(OAIAPIKey)

	default:
		var OHostPort string

		OHost := os.Getenv("OLLAMA_HOST")
		OHostPath := os.Getenv("OLLAMA_HOST_PATH")

		if len(OHostPath) <= 0 {
			OHostPath = "/"
		}

		OHostSchema, OHostURL, isOK := strings.Cut(OHost, "://")

		if !isOK {
			OHostSchema = "http"
			OHostURL = OHost
			OHostPort = "11434"
		}

		if OHostPort == "" {
			switch OHostSchema {
			case "http":
				OHostPort = "80"

			case "https":
				OHostPort = "443"
			}
		}

		OClient = Ollama.NewClient(&url.URL{
			Scheme: OHostSchema,
			Host:   net.JoinHostPort(OHostURL, OHostPort),
			Path:   OHostPath,
		}, http.DefaultClient)
	}
}

func GPT3Completion(question string) (err error) {
	switch strings.ToLower(GPTEngine) {
	case "openai":
		OAIGPTRequest := OpenAI.ChatCompletionRequest{
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

		OAIGPTResponse, err := OAIClient.CreateChatCompletion(
			context.Background(),
			OAIGPTRequest,
		)

		if err != nil {
			return err
		}

		if len(OAIGPTResponse.Choices) > 0 {
			fmt.Println(OAIGPTResponse.Choices[0].Message.Content)
		}

		return nil

	default:
		var OGPTResponseText string

		isStream := new(bool)
		*isStream = false

		OGPTRequest := &Ollama.ChatRequest{
			Model: "dadang",
			Messages: []Ollama.Message{
				{
					Role:    "user",
					Content: question,
				},
			},
			Stream: isStream,
		}

		OGTPResponseFunc := func(OGPTResponse Ollama.ChatResponse) error {
			OGPTResponseText = OGPTResponse.Message.Content
			return nil
		}

		err := OClient.Chat(
			context.Background(),
			OGPTRequest,
			OGTPResponseFunc,
		)

		if err != nil {
			return err
		}

		if len(OGPTResponseText) > 0 {
			fmt.Println(OGPTResponseText)
		}

		return nil

	}
}
