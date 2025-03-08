package app

import (
	"fmt"
	"os"

	"github.com/reinhardlinardi/atm-report/internal/config/filecron"

	"gopkg.in/yaml.v3"
)

type Config struct {
	DB       *DBConfig        `yaml:"db"`
	FileCron *filecron.Config `yaml:"filecron"`
}

type DBConfig struct {
	User   string `yaml:"user"`
	Pass   string `yaml:"pass"`
	Host   string `yaml:"host"`
	Port   uint16 `yaml:"port"`
	Schema string `yaml:"schema"`
}

func Parse(path string) (*Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &Config{}

	if err := yaml.Unmarshal(b, config); err != nil {
		fmt.Printf("err yaml unmarshal: %s\n", err.Error())
		return nil, err
	}

	if err := parseDB(config.DB); err != nil {
		fmt.Printf("err parse db config: %s\n", err.Error())
		return nil, err
	}
	if err := filecron.Parse(config.FileCron); err != nil {
		fmt.Printf("err parse filecron config: %s\n", err.Error())
		return nil, err
	}

	return config, nil
}

func parseDB(config *DBConfig) error {
	return nil
}
