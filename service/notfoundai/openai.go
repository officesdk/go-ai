package openai

import (
	"context"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/officesdk/go-ai/aimodel"
	"github.com/officesdk/go-ai/config"
	"github.com/officesdk/go-ai/manager"
)

type NotFoundAIService struct {
	client *resty.Client
	config config.Config
}

var Service = &NotFoundAIService{}

var _ manager.AIService = (*NotFoundAIService)(nil)

func init() {
	manager.RegisterAIService(Service)
}

func (o NotFoundAIService) Init(config config.Config) error {
	o.config = config
	return nil
}

func (o NotFoundAIService) Name() string {
	return "notfound"
}

func (o NotFoundAIService) ChatCompletion(ctx context.Context, request aimodel.ChatCompletionRequest) (response aimodel.ChatCompletionResponse, err error) {
	return aimodel.ChatCompletionResponse{}, fmt.Errorf("not found ai")
}

func (o NotFoundAIService) ChatCompletionStream(ctx context.Context, request aimodel.ChatCompletionRequest) (response *aimodel.ChatCompletionStream, err error) {
	return &aimodel.ChatCompletionStream{}, fmt.Errorf("not found ai")
}
