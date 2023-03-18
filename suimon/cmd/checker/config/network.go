package config

import (
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/pkg/env"
	"github.com/bartosian/sui_helpers/suimon/internal/pkg/log"
)

const (
	networkNameDefault         = "devnet"
	invalidNetworkTypeProvided = "provide valid network type by using -n option or set it in suimon.yaml"
)

// ParseNetworkConfig decodes the network configuration for the given network name and maps the decoded data into a SuimonConfig struct.
// This function accepts the following parameters:
// - suimonConfig: a pointer to a SuimonConfig struct that will be populated with the network configuration data.
// - networkName: a pointer to a string representing the name of the network to be configured.
// The function returns an enums.NetworkType representing the type of the network, and an error if there was an issue parsing the configuration data.
func ParseNetworkConfig(suimonConfig *SuimonConfig, networkName *string) (enums.NetworkType, error) {
	logger := log.NewLogger()

	networkConfig := suimonConfig.Network

	if *networkName == "" && networkConfig.Name == "" {
		envValue := env.GetEnvWithDefault("SUIMON_NETWORK", networkNameDefault)

		networkName = &envValue
	}

	if *networkName == "" && networkConfig.Name != "" {
		networkName = &networkConfig.Name
	}

	result, err := enums.NetworkTypeFromString(*networkName)
	if err != nil {
		logger.Error(invalidNetworkTypeProvided)
	}

	return result, err
}
