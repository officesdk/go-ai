# Go-AI
<p align="center">
  <a href="https://github.com/officesdk/go-ai/releases"><img src="https://img.shields.io/github/v/release/officesdk/go-ai?style=flat" alt="Release"></a>
  <a href="https://github.com/officesdk/go-ai/stargazers"><img src="https://img.shields.io/github/stars/officesdk/go-ai?style=flat" alt="Stars"></a>
  <a href="https://github.com/officesdk/go-ai/network/members"><img src="https://img.shields.io/github/forks/officesdk/go-ai?style=flat" alt="Forks"></a>
  <a href="https://github.com/officesdk/go-ai/issues"><img src="https://img.shields.io/github/issues/officesdk/go-ai?color=gold&style=flat" alt="Issues"></a>
  <a href="https://github.com/officesdk/go-ai/pulls"><img src="https://img.shields.io/github/issues-pr/officesdk/go-ai?color=gold&style=flat" alt="Pull Requests"></a>
  <a href="https://github.com/officesdk/go-ai/blob/main/LICENSE"><img src="https://img.shields.io/badge/license-MIT-green.svg" alt="License"></a>
  <a href="https://github.com/officesdk/go-ai/graphs/contributors"><img src="https://img.shields.io/github/contributors/officesdk/go-ai?color=green&style=flat" alt="Contributors"></a>
  <a href="https://github.com/officesdk/go-ai/commits"><img src="https://img.shields.io/github/last-commit/officesdk/go-ai?color=green&style=flat" alt="Last Commit"></a>
</p>
<p align="center">
  <a href="https://pkg.go.dev/github.com/officesdk/go-ai"><img src="https://img.shields.io/badge/-reference-blue?logo=go&logoColor=white&style=flat" alt="Go Reference"></a>
  <a href="https://goreportcard.com/report/github.com/officesdk/go-ai"><img src="https://img.shields.io/badge/go%20report-A+-brightgreen?style=flat" alt="Go Report"></a>
  <a href="https://github.com/officesdk/go-ai/actions"><img src="https://img.shields.io/badge/Go%20Tests-passing-brightgreen?style=flat" alt="Go Tests"></a>
</p>

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
	client,err := goai.NewClient(goai.WithConfig(config.Config{
		Name: "openai",
		ApiHost: "your host",
		ApiKey: "your key",
		Enabled: true,
		ApiTimeout: 120 * time.Second,
    }))
	
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
	client,err := goai.NewClient(goai.WithConfig(config.Config{
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
	}))
	
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

### Multi AI example use json config:

```go
package main

import (
	"context"
	"fmt"
	goai "github.com/officesdk/go-ai"
	"github.com/officesdk/go-ai/config"
)

func main() {
	client,err := goai.NewClient(goai.WithRawConfig([]byte(`
		[{
			"name": "openai",
			"apiHost": "your host",
			"apiKey": "your key",
			"enabled": true,
			"apiTimeout": "120s"	
		}]
	`)))
	
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