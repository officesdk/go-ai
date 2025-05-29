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

func WithRawConfig(rawConfig []byte) Option {
	return func(c *Client) {
		c.rawConfig = rawConfig
	}
}

func WithRawConfigType(rawConfigType config.ConfigType) Option {
	return func(c *Client) {
		c.rawConfigType = rawConfigType
	}
}
