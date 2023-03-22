package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/bartosian/sui_helpers/suimon/internal/pkg/env"
	"github.com/bartosian/sui_helpers/suimon/internal/pkg/log"
)

const (
	nodeConfigPath         = "%s/.suimon/fullnode.yaml"
	configFullnodeNotFound = "provide path to the fullnode.yaml file by using -f option or by setting SUIMON_NODE_CONFIG_PATH env variable or set path to this file in suimon.yaml"
	configFullnodeInvalid  = "make sure fullnode.yaml file has correct syntax and properties"
)

type NodeConfig struct {
	DbPath                string `yaml:"db-path"`
	NetworkAddress        string `yaml:"network-address"`
	MetricsAddress        string `yaml:"metrics-address"`
	JSONRPCAddress        string `yaml:"json-rpc-address"`
	WebsocketAddress      string `yaml:"websocket-address"`
	EnableEventProcessing bool   `yaml:"enable-event-processing"`
	ConsensusConfig       any    `yaml:"consensus-config"`
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
	logger := log.NewLogger()
	configPath := *path

	if configPath == "" {
		home := os.Getenv("HOME")
		configPath = env.GetEnvWithDefault("SUIMON_NODE_CONFIG_PATH", fmt.Sprintf(nodeConfigPath, home))
	}

	file, err := os.ReadFile(configPath)
	if err != nil && suimonPath == "" {
		logger.Error(configFullnodeNotFound)

		return nil, err
	}

	file, err = os.ReadFile(suimonPath)
	if err != nil {
		logger.Error(configFullnodeNotFound)

		return nil, err
	}

	var result NodeConfig
	err = yaml.Unmarshal(file, &result)
	if err != nil {
		logger.Error(configFullnodeInvalid)

		return nil, err
	}

	return &result, nil
}
