package go_ai

import (
	"encoding/json"
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/mitchellh/mapstructure"
	"github.com/officesdk/go-ai/config"
	"github.com/officesdk/go-ai/manager"
	_ "github.com/officesdk/go-ai/service/notfoundai" // Import the not found AI service
	_ "github.com/officesdk/go-ai/service/openai"     // Import the not found AI service
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
		configArr := make([]config.Config, 0)
		tagName := "json"
		storeMap := make([]any, 0)
		switch c.rawConfigType {
		case config.ConfigTypeJson:
			// JSON configuration
			err := json.Unmarshal(c.rawConfig, &storeMap)
			if err != nil {
				return nil, fmt.Errorf("failed to json unmarshal raw config: %w", err)
			}
		case config.ConfigTypeYaml:
			// YAML configuration
			tagName = "yaml"
			err := yaml.Unmarshal(c.rawConfig, &storeMap)
			if err != nil {
				return nil, fmt.Errorf("failed to yaml unmarshal raw config: %w", err)
			}
		case config.ConfigTypeToml:
			// TOML configuration
			tagName = "toml"
			err := toml.Unmarshal(c.rawConfig, &storeMap)
			if err != nil {
				return nil, fmt.Errorf("failed to toml unmarshal raw config: %w", err)
			}
		default:
			return nil, fmt.Errorf("unsupported config type: %s", c.rawConfigType)
		}

		decoderConfig := mapstructure.DecoderConfig{
			DecodeHook:       mapstructure.StringToTimeDurationHookFunc(),
			Result:           &configArr,
			TagName:          tagName,
			WeaklyTypedInput: false,
			Squash:           false,
		}
		decoder, err := mapstructure.NewDecoder(&decoderConfig)
		if err != nil {
			return nil, err
		}

		err = decoder.Decode(storeMap)
		if err != nil {
			return nil, err
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
