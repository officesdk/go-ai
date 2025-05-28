package go_ai

import (
	"encoding/json"
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/officesdk/go-ai/config"
	"github.com/officesdk/go-ai/manager"
	"gopkg.in/yaml.v3"
)

type Client struct {
	rawConfig     []byte
	rawConfigType config.ConfigType
	config        []config.Config
}

func NewClient(options ...Option) (*Client, error) {
	c := &Client{
		rawConfigType: config.ConfigTypeJson,
	}
	for _, option := range options {
		option(c)
	}
	if len(c.rawConfig) != 0 {
		var configArr []config.Config
		switch c.rawConfigType {
		case config.ConfigTypeJson:
			// Parse JSON configuration

			err := json.Unmarshal(c.rawConfig, &configArr)
			if err != nil {
				return nil, fmt.Errorf("failed to parse JSON config: %w", err)
			}
		case config.ConfigTypeYaml:
			// Parse YAML configuration
			err := yaml.Unmarshal(c.rawConfig, &configArr)
			if err != nil {
				return nil, fmt.Errorf("failed to parse YAML config: %w", err)
			}
		case config.ConfigTypeToml:
			// Parse TOML configuration
			err := toml.Unmarshal(c.rawConfig, &configArr)
			if err != nil {
				return nil, fmt.Errorf("failed to parse TOML config: %w", err)
			}
		default:
			return nil, fmt.Errorf("unsupported config type: %s", c.rawConfigType)
		}
		configMap := make(map[string]config.Config)
		for _, cfg := range c.config {
			configMap[cfg.Name] = cfg
		}
		for _, cfg := range configArr {
			configMap[cfg.Name] = cfg
		}
		// 清空config
		c.config = make([]config.Config, 0, len(configMap))
		// 重新赋值
		for _, cfg := range configMap {
			c.config = append(c.config, cfg)
		}
	}

	// 初始化服务
	for _, cfg := range c.config {
		aiService, flag := manager.GetAIService(cfg.Name)
		if !flag {
			return nil, fmt.Errorf("service %s not found", cfg.Name)
		}

		if err := aiService.Init(cfg); err != nil {
			return nil, fmt.Errorf("failed to init config %s: %w", cfg.Name, err)
		}
	}
	return c, nil
}

func (c *Client) Use(name string) manager.AIService {
	if service, ok := manager.GetAIService(name); ok {
		return service
	}
	notfoundSvc, _ := manager.GetAIService("notfound")
	return notfoundSvc
}
