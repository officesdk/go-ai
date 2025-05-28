package go_ai

import (
	"github.com/officesdk/go-ai/config"
)

// Option overrides a Client's default configuration.
type Option func(c *Client)

func WithConfig(options ...config.Config) Option {
	return func(c *Client) {
		c.config = append(c.config, options...)
	}
}
