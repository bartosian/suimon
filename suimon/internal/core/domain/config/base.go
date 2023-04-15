package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/bartosian/sui_helpers/suimon/internal/pkg/env"
)

const (
	suimonConfigPath   = "%s/.suimon/suimon.yaml"
	suimonConfigEnvVar = "SUIMON_CONFIG_PATH"
)

type Config struct {
	PublicRPC []string `yaml:"public-rpc"`
	FullNodes []struct {
		JSONRPCAddress string `yaml:"json-rpc-address"`
		MetricsAddress string `yaml:"metrics-address"`
	} `yaml:"full-nodes"`
	Validators []struct {
		MetricsAddress string `yaml:"metrics-address"`
	} `yaml:"validators"`
	SeedPeers []string `yaml:"seed-peers"`
	IPLookup  struct {
		AccessToken string `yaml:"access-token"`
	} `yaml:"ip-lookup"`
}

func NewConfig() (*Config, error) {
	home := os.Getenv("HOME")
	configPath := env.GetEnvWithDefault(suimonConfigEnvVar, fmt.Sprintf(suimonConfigPath, home))

	file, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var result Config
	err = yaml.Unmarshal(file, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
