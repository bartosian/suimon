package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/bartosian/sui_helpers/suimon/internal/pkg/env"
	"github.com/bartosian/sui_helpers/suimon/internal/pkg/log"
)

const (
	suimonConfigPath   = "%s/.suimon/suimon.yaml"
	suimonConfigEnvVar = "SUIMON_CONFIG_PATH"
)

type Config struct {
	MonitorsConfig struct {
		RPCTable struct {
			Display bool `yaml:"display"`
		} `yaml:"rpcgw-table"`
		NodeTable struct {
			Display bool `yaml:"display"`
		} `yaml:"node-table"`
		ValidatorTable struct {
			Display bool `yaml:"display"`
		} `yaml:"validator-table"`
		PeersTable struct {
			Display bool `yaml:"display"`
		} `yaml:"peers-table"`
		SystemStateTable struct {
			Display bool `yaml:"display"`
		} `yaml:"system-state-table"`
		ValidatorsCountsTable struct {
			Display bool `yaml:"display"`
		} `yaml:"validators-counts-table"`
		ValidatorsAtRiskTable struct {
			Display bool `yaml:"display"`
		} `yaml:"validators-at-risk-table"`
		ValidatorReportsTable struct {
			Display bool `yaml:"display"`
		} `yaml:"validator-reports-table"`
		ActiveValidatorsTable struct {
			Display bool `yaml:"display"`
		} `yaml:"active-validators-table"`
	} `yaml:"monitors-tables"`
	PublicRPC []string `yaml:"public-rpcgw"`
	FullNode  struct {
		JSONRPCAddress string `yaml:"json-rpcgw-address"`
		MetricsAddress string `yaml:"rpcgw-address"`
	} `yaml:"full-node"`
	Validator struct {
		MetricsAddress string `yaml:"rpcgw-address"`
	} `yaml:"validator"`
	SeedPeers []string `yaml:"seed-peers"`
	IPLookup  struct {
		AccessToken string `yaml:"access-token"`
	} `yaml:"ip-lookup"`
	MonitorsVisual struct {
		EnableEmojis bool `yaml:"enable-emojis"`
	} `yaml:"monitors-visual"`
}

func NewConfig(logger log.Logger) (*Config, error) {
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
