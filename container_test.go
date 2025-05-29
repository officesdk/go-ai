package go_ai

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/officesdk/go-ai/aimodel"
	"github.com/officesdk/go-ai/config"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	got, err := NewClient(WithConfig(config.Config{
		Name:       "openai",
		ApiHost:    "your host",
		ApiKey:     "your key",
		Enabled:    true,
		ApiTimeout: 120 * time.Second,
	}))
	assert.NoError(t, err)
	assert.Equal(t, 1, len(got.config))
	assert.NotNil(t, got)
	assert.Equal(t, "openai", got.config[0].Name)
	assert.Equal(t, "your host", got.config[0].ApiHost)
	assert.Equal(t, "your key", got.config[0].ApiKey)
	assert.Equal(t, true, got.config[0].Enabled)
	assert.Equal(t, 120*time.Second, got.config[0].ApiTimeout)
}

func TestNewClientRawConfig(t *testing.T) {
	got, err := NewClient(WithRawConfig([]byte(`
		[{
			"name": "openai",
			"apiHost": "your host",
			"apiKey": "your key",
			"enabled": true,
			"apiTimeout": "120s"	
		}]
	`)))
	assert.NoError(t, err)
	assert.Equal(t, 1, len(got.config))
	assert.NotNil(t, got)
	assert.Equal(t, "openai", got.config[0].Name)
	assert.Equal(t, "your host", got.config[0].ApiHost)
	assert.Equal(t, "your key", got.config[0].ApiKey)
	assert.Equal(t, true, got.config[0].Enabled)
	assert.Equal(t, 120*time.Second, got.config[0].ApiTimeout)
}

func TestNewClientChat(t *testing.T) {
	client, err := NewClient(WithConfig(config.Config{
		Name:       "openai",
		ApiKey:     os.Getenv("API_KEY"),
		ApiHost:    os.Getenv("API_HOST"),
		Enabled:    true,
		ApiTimeout: 120 * time.Second,
	}))
	assert.NoError(t, err)
	gotResponse, err := client.Use("openai").ChatCompletion(context.Background(), aimodel.ChatCompletionRequest{
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
