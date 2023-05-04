package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/bartosian/suimon/cmd/checker/enums"
	"github.com/bartosian/suimon/pkg/env"
	"github.com/bartosian/suimon/pkg/log"
)

const (
	suimonConfigPath     = "%s/.suimon/suimon.yaml"
	configSuimonNotFound = "provide path to the suimon.yaml file by using -s option or by setting SUIMON_CONFIG_PATH env variable or put suimon.yaml in $HOME/.suimon/suimon.yaml"
	configSuimonInvalid  = "make sure suimon.yaml file has correct syntax and properties"
)

type SuimonConfig struct {
	MonitorsConfig struct {
		RPCTable struct {
			Display bool `yaml:"display"`
		} `yaml:"rpc-table"`
		NodeTable struct {
			Display bool `yaml:"display"`
		} `yaml:"node-table"`
		PeersTable struct {
			Display bool `yaml:"display"`
		} `yaml:"peers-table"`
	} `yaml:"monitors-config"`
	RPCConfig struct {
		Testnet []string `yaml:"testnet"`
		Devnet  []string `yaml:"devnet"`
	} `yaml:"rpc-config"`
	NodeConfigPath string `yaml:"node-config-path"`
	Network        string `yaml:"network"`
	NetworkType    enums.NetworkType
	IPLookup       struct {
		AccessToken string `yaml:"access-token"`
	} `yaml:"ip-lookup"`
	MonitorsVisual struct {
		ColorScheme  string `yaml:"color-scheme"`
		EnableEmojis bool   `yaml:"enable-emojis"`
	} `yaml:"monitors-visual"`
}

func ParseSuimonConfig(path *string) (*SuimonConfig, error) {
	logger := log.NewLogger()
	configPath := *path

	if configPath == "" {
		home := os.Getenv("HOME")
		configPath = env.GetEnvWithDefault("SUIMON_CONFIG_PATH", fmt.Sprintf(suimonConfigPath, home))
	}

	file, err := os.ReadFile(configPath)
	if err != nil {
		logger.Error(configSuimonNotFound)

		return nil, err
	}

	var result SuimonConfig
	err = yaml.Unmarshal(file, &result)
	if err != nil {
		logger.Error(configSuimonInvalid)

		return nil, err
	}

	return &result, nil
}

func (sconfig *SuimonConfig) SetNetworkConfig(network enums.NetworkType) {
	sconfig.NetworkType = network
	sconfig.Network = network.String()
}

func (sconfig *SuimonConfig) GetRPCByNetwork() []string {
	switch sconfig.NetworkType {
	case enums.NetworkTypeDevnet:
		return sconfig.RPCConfig.Devnet
	case enums.NetworkTypeTestnet:
		return sconfig.RPCConfig.Testnet
	}

	return nil
}
