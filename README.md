# go-ai
go ai sdk


## Usage
### ChatGPT example usage:

```go
package main

import (
	"context"
	"fmt"
	goai "github.com/officesdk/go-ai"
	"github.com/officesdk/go-ai/config"
)

func main() {
	client,err := goai.NewClient(config.Config{
		ApiHost: "your host",
		ApiKey: "your key",
		Enabled: true,
		ApiTimeout: 120 * time.Second,
    })
	
	resp, err := client.Use("openai").ChatCompletion(
		context.Background(),
		aimodel.ChatCompletionRequest{
			Model: "gpt-3.5-turbo",
			Messages: []aimodel.ChatCompletionMessage{
				{
					Role:    aimodel.ChatMessageRoleUser,
					Content: "Hello!",
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}

```