package config

import (
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"
	"github.com/bartosian/sui_helpers/suimon/pkg/env"
)

const (
	networkConfig = "devnet"
)

func ParseNetworkConfig(suimonConfig *SuimonConfig, network *string) (enums.NetworkType, error) {
	if *network == "" && suimonConfig.Network == "" {
		envValue := env.GetEnvWithDefault("SUIMON_NETWORK", networkConfig)

		network = &envValue
	}

	if *network == "" && suimonConfig.Network != "" {
		network = &suimonConfig.Network
	}

	return enums.NetworkTypeFromString(*network)
}
