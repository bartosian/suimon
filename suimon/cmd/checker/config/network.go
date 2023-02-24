package config

import (
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"
	"github.com/bartosian/sui_helpers/suimon/pkg/env"
	"github.com/bartosian/sui_helpers/suimon/pkg/log"
)

const (
	networkConfig              = "devnet"
	invalidNetworkTypeProvided = "provide valid network type by using -n option or set it in suimon.yaml"
)

func ParseNetworkConfig(suimonConfig *SuimonConfig, network *string) (enums.NetworkType, error) {
	logger := log.NewLogger()

	if *network == "" && suimonConfig.Network == "" {
		envValue := env.GetEnvWithDefault("SUIMON_NETWORK", networkConfig)

		network = &envValue
	}

	if *network == "" && suimonConfig.Network != "" {
		network = &suimonConfig.Network
	}

	result, err := enums.NetworkTypeFromString(*network)
	if err != nil {
		logger.Error(invalidNetworkTypeProvided)
	}

	return result, err
}
