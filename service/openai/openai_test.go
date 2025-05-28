package openai

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/officesdk/go-ai/aimodel"
	"github.com/officesdk/go-ai/config"
	"github.com/stretchr/testify/assert"
)

func TestOpenAIService_ChatCompletion(t *testing.T) {
	o := &OpenAIService{}
	o.Init(config.Config{
		Name:       "openai",
		ApiKey:     os.Getenv("API_KEY"),
		ApiHost:    os.Getenv("API_HOST"),
		ApiTimeout: 120 * time.Second,
		Debug:      true,
		Enabled:    true,
	})

	gotResponse, err := o.ChatCompletion(context.Background(), aimodel.ChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []aimodel.ChatCompletionMessage{
			{
				Role:    aimodel.ChatMessageRoleUser,
				Content: "Hello!",
			},
		},
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, gotResponse)
}
