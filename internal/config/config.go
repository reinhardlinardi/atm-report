package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server *ServerConfig `yaml:"server"`
	DB     *DBConfig     `yaml:"db"`
	Cron   *CronConfig   `yaml:"cron"`
}

type ServerConfig struct {
	Port uint16 `yaml:"port"`
}

type DBConfig struct {
	User   string `yaml:"user"`
	Pass   string `yaml:"pass"`
	Host   string `yaml:"host"`
	Port   uint16 `yaml:"port"`
	Schema string `yaml:"schema"`
}

type CronConfig struct {
	Path string `yaml:"path"`
}

func Parse(path string) (*Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &Config{}

	if err := yaml.Unmarshal(b, config); err != nil {
		return nil, err
	}

	return config, nil
}
