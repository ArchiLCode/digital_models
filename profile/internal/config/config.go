package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"os"
	"time"
)

type Config struct {
	Env         string     `yaml:"env"`
	StoragePath string     `yaml:"storage_path" env-required:"true"`
	GRPC        GRPCConfig `yaml:"grpc"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("config file path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file not found")
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config" + err.Error())
	}
	return &cfg
}

func MustLoadWithPath(path string) *Config {
	if path == "" {
		panic("test config file path is empty")
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("test config file not found")
	}
	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read test config: " + err.Error())
	}
	return &cfg
}

func fetchConfigPath() string {
	err := godotenv.Load("/app/.env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	return os.Getenv("CONFIG_PATH")
}
