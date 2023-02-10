package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/bartosian/sui_helpers/suimon/pkg/env"
)

const (
	nodeConfigPath = "%s/.suimon/fullnode.yaml"
)

type NodeConfig struct {
	DbPath                string `yaml:"db-path"`
	NetworkAddress        string `yaml:"network-address"`
	MetricsAddress        string `yaml:"metrics-address"`
	JSONRPCAddress        string `yaml:"json-rpc-address"`
	WebsocketAddress      string `yaml:"websocket-address"`
	EnableEventProcessing bool   `yaml:"enable-event-processing"`
	Genesis               struct {
		GenesisFileLocation string `yaml:"genesis-file-location"`
	} `yaml:"genesis"`
	P2PConfig struct {
		SeedPeers []struct {
			Address string `yaml:"address"`
		} `yaml:"seed-peers"`
	} `yaml:"p2p-config"`
}

func ParseNodeConfig(path *string, suimonPath string) (*NodeConfig, error) {
	configPath := *path

	if configPath == "" {
		home := os.Getenv("HOME")
		configPath = env.GetEnvWithDefault("SUIMON_NODE_CONFIG_PATH", fmt.Sprintf(nodeConfigPath, home))
	}

	configPath = "newfile.sh"

	file, err := os.ReadFile(configPath)
	if err != nil && suimonPath == "" {
		return nil, err
	}

	file, err = os.ReadFile(suimonPath)
	if err != nil {
		return nil, err
	}

	var result NodeConfig
	err = yaml.Unmarshal(file, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
