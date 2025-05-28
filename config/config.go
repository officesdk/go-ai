package config

import (
	"time"
)

type Config struct {
	Name       string
	ApiKey     string
	ApiHost    string
	ApiTimeout time.Duration
	Debug      bool
	Enabled    bool
}

type ConfigType string

const (
	ConfigTypeJson ConfigType = "json"
	ConfigTypeYaml ConfigType = "yaml"
	ConfigTypeToml ConfigType = "toml"
)
