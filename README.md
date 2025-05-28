# go-ai
go ai sdk


## Usage
### Single AI ChatGPT example usage:

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
		Name: "openai",
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

### Multi AI example usage:

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
		Name: "openai",
		ApiHost: "your host",
		ApiKey: "your key",
		Enabled: true,
		ApiTimeout: 120 * time.Second,
    },config.Config{
		Name: "deepseek",
		ApiHost: "your host",
		ApiKey: "your key",
		Enabled: true,
		ApiTimeout: 120 * time.Second,
	})
	
	resp, err := client.Use("deepseek").ChatCompletion(
		context.Background(),
		aimodel.ChatCompletionRequest{
			Model: "deepseek-r1",
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