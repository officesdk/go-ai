package manager

import (
	"context"

	"github.com/officesdk/go-ai/aimodel"
	"github.com/officesdk/go-ai/config"
)

// AIService is the interface for ai service
type AIService interface {
	Init(config config.Config) error
	Name() string
	ChatCompletion(ctx context.Context, request aimodel.ChatCompletionRequest) (response aimodel.ChatCompletionResponse, err error)
	ChatCompletionStream(ctx context.Context, request aimodel.ChatCompletionRequest) (response aimodel.ChatCompletionResponse, err error)
}
