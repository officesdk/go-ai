package openai

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/gotomicro/ego/client/ehttp"
	"github.com/officesdk/go-ai/aimodel"
	"github.com/officesdk/go-ai/config"
	"github.com/officesdk/go-ai/manager"
)

type OpenAIService struct {
	client *resty.Client
	config config.Config
}

var Service = &OpenAIService{}

var _ manager.AIService = (*OpenAIService)(nil)

func init() {
	manager.RegisterAIService(Service)
}

func (o *OpenAIService) Init(config config.Config) error {
	o.config = config
	o.client = ehttp.DefaultContainer().Build().SetDebug(o.config.Debug).SetBaseURL(o.config.ApiHost).SetTimeout(o.config.ApiTimeout)
	return nil
}

func (o *OpenAIService) Name() string {
	return "openai"
}

func (o *OpenAIService) ChatCompletion(ctx context.Context, request aimodel.ChatCompletionRequest) (response aimodel.ChatCompletionResponse, err error) {
	request.Stream = false
	resp, err := o.client.R().SetContext(ctx).SetHeaders(map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + o.config.ApiKey,
	}).SetBody(request).Post("/v1/chat/completions")
	if err != nil {
		if strings.Contains(err.Error(), "connection refused") {
			err = fmt.Errorf("post fail, err: %w", err)
			return
		}
		err = fmt.Errorf("error sending request: %w", err)
		return
	}

	if resp.StatusCode() >= 400 {
		err = fmt.Errorf("unexpected status code: %d", resp.StatusCode())
		return
	}
	err = json.Unmarshal(resp.Body(), &response)
	if err != nil {
		err = fmt.Errorf("error decoding response: %w", err)
		return
	}
	return
}

// CreateChatCompletionStream — API call to create a chat completion w/ streaming
// support. It sets whether to stream back partial progress. If set, tokens will be
// sent as data-only server-sent events as they become available, with the
// stream terminated by a data: [DONE] message.
func (c *OpenAIService) ChatCompletionStream(
	ctx context.Context,
	request aimodel.ChatCompletionRequest,
) (stream *aimodel.ChatCompletionStream, err error) {
	request.Stream = true

	resp, err := c.client.R().SetContext(ctx).SetDoNotParseResponse(true).SetHeaders(map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + c.config.ApiKey,
		"Cache-Control": "no-cache",
		"Connection":    "keep-alive",
		"Accept":        "text/event-stream",
	}).SetBody(request).Post("/v1/chat/completions")
	if err != nil {
		if strings.Contains(err.Error(), "connection refused") {
			return
		}
		return
	}

	streamReader := &aimodel.StreamReader{
		Reader: bufio.NewReader(resp.RawResponse.Body),
	}
	return &aimodel.ChatCompletionStream{
		StreamReader: streamReader,
	}, nil
}
