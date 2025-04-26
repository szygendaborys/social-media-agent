package ai

import (
	"github.com/openai/openai-go"
)

type Conversation struct {
	client OpenAIClient
	model  string
}

func InitConversation() Conversation {
	client := CreateOpenAIClient()

	return Conversation{
		client: client,
		model:  "smollm",
	}
}

func (c *Conversation) askAI(q string) (string, error) {
	println("Calling OpenAI to ask about " + q)

	messages := []openai.ChatCompletionMessageParamUnion{
		openai.UserMessage(q),
	}

	return SendChatRequest(&c.client, openai.ChatCompletionNewParams{
		Messages: messages,
		Model:    c.model,
	})
}
