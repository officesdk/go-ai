package config

import (
	"time"
)

type Config struct {
	Name       string        `json:"name"`
	ApiKey     string        `json:"apiKey"`
	ApiHost    string        `json:"apiHost"`
	ApiTimeout time.Duration `json:"apiTimeout"`
	Debug      bool          `json:"debug"`
	Enabled    bool          `json:"enabled"`
}

type ConfigType string

const (
	ConfigTypeJson ConfigType = "json"
	ConfigTypeYaml ConfigType = "yaml"
	ConfigTypeToml ConfigType = "toml"
)
