package appConfig

import (
	"os"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	Server ServerConfig `yaml:"server"`
	Worker WorkerConfig `yaml:"worker"`
}

type ServerConfig struct {
	Port int      `yaml:"port"`
	DB   DBConfig `yaml:"db"`
}

type DBConfig struct {
	DSN string `yaml:"dsn,omitempty"`
}

type WorkerConfig struct {
	PollInterval int `yaml:"poll-interval"`
	JobTimeout   int `yaml:"job-timeout"`
	RetryBackoff int `yaml:"retry-backoff"`
	MaxRetries   int `yaml:"max-retries"`
	Concurrency  int `yaml:"concurrency"`
}

// func DefaultAppConfig() *AppConfig {
// 	return &AppConfig{
// 		Server: ServerConfig{
// 			Port: 8080,
// 			DB: DBConfig{
// 				DSN: "default vaoue",
// 			},
// 		},
// 	}
// }

func LoadConfig(path string) (*AppConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// this will replace ${VAR} with actual values from the env
	expanded := os.ExpandEnv(string(data))

	var cfg AppConfig
	if err := yaml.Unmarshal([]byte(expanded), &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
