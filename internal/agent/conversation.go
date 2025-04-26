package agent

import (
	"github.com/openai/openai-go"
	"github.com/szygendaborys/social-media-agent/internal/ai"
)

type Conversation struct {
	client ai.OpenAIClient
	model  string
}

func InitConversation() Conversation {
	client := ai.CreateOpenAIClient()

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

	return ai.AskAI(&c.client, openai.ChatCompletionNewParams{
		Messages: messages,
		Model:    c.model,
	})
}
