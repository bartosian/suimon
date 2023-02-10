package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"
	"github.com/bartosian/sui_helpers/suimon/pkg/env"
)

const (
	suimonConfigPath = "%s/.suimon/suimon.yaml"
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
	HostLookupConfig struct {
		EnableLookup bool   `yaml:"enable-lookup"`
		GeoDbPath    string `yaml:"geo-db-path"`
	} `yaml:"host-lookup-config"`
	NodeConfigPath string `yaml:"node-config-path"`
	Network        string `yaml:"network"`
	NetworkType    enums.NetworkType
}

func ParseSuimonConfig(path *string) (*SuimonConfig, error) {
	configPath := *path

	if configPath == "" {
		home := os.Getenv("HOME")
		configPath = env.GetEnvWithDefault("SUIMON_CONFIG_PATH", fmt.Sprintf(suimonConfigPath, home))
	}

	file, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var result SuimonConfig
	err = yaml.Unmarshal(file, &result)
	if err != nil {
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
