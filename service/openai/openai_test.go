package openai

import (
	"context"
	"errors"
	"fmt"
	"io"
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

func TestOpenAIService_ChatCompletionStream(t *testing.T) {
	o := &OpenAIService{}
	o.Init(config.Config{
		Name:       "openai",
		ApiKey:     os.Getenv("API_KEY"),
		ApiHost:    os.Getenv("API_HOST"),
		ApiTimeout: 120 * time.Second,
		Debug:      true,
		Enabled:    true,
	})

	gotResponse, err := o.ChatCompletionStream(context.Background(), aimodel.ChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []aimodel.ChatCompletionMessage{
			{
				Role:    aimodel.ChatMessageRoleUser,
				Content: "Hello, 如何写好一个Golang!",
			},
		},
	})
	assert.NoError(t, err)

	for {
		var response aimodel.ChatCompletionStreamResponse
		response, err = gotResponse.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("Stream finished")
			return
		}

		if err != nil {
			fmt.Printf("Stream error: %v\n", err)
			return
		}

		fmt.Printf("Stream response: %#v\n", response)
	}

}
