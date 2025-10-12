package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string        `yaml:"env" env-default:"local"`
	StoragePath string        `yaml:"storage_path" env-required:"true"`
	TokenTTL    time.Duration `yaml:"token_ttl" env-required:"1h"`
	GRPC        GRPCConfig    `yaml:"grpc"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timout"`
}

func MustLoad() *Config {
	configPath := FetchConfigPath()

	if configPath == "" {
		panic("path is empty")
	}

	return MustLoadPath(configPath)

}

func MustLoadPath(configPath string) *Config {

	if _, err := os.Stat(configPath); err != nil {
		panic("file does not exist" + err.Error())
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("failed to read config" + err.Error())
	}
	return &cfg

}

func FetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res

}
