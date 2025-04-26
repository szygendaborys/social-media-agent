package ai

import (
	"context"
	"os"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type OpenAIClient struct {
	client openai.Client
}

func CreateOpenAIClient() OpenAIClient {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		panic("OPENAI_API_KEY is not defined")
	}

	baseURL := os.Getenv("OPENAI_BASE_URL")
	if baseURL == "" {
		panic("OPENAI_BASE_URL is not defined")
	}

	return OpenAIClient{
		client: openai.NewClient(
			option.WithAPIKey(apiKey),
			option.WithBaseURL(baseURL),
		),
	}
}

func SendChatRequest(c *OpenAIClient, params openai.ChatCompletionNewParams) (string, error) {
	chatCompletion, err := c.client.Chat.Completions.New(context.TODO(), params)

	if err != nil {
		return "", err
	}

	return chatCompletion.Choices[0].Message.Content, nil
}
