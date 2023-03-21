package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"

	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/pkg/env"
	"github.com/bartosian/sui_helpers/suimon/internal/pkg/log"
)

const (
	defaultServiceName = "suid"
	defaultImageName   = "sui-node"
	defaultScreenName  = "sui"

	suimonConfigPath     = "%s/.suimon/suimon.yaml"
	configSuimonNotFound = "provide path to the suimon.yaml file by using -s option or by setting SUIMON_CONFIG_PATH env variable or put suimon.yaml in $HOME/.suimon/suimon.yaml"
	configSuimonInvalid  = "make sure suimon.yaml file has correct syntax and properties"
)

type (
	ProcessLaunchType struct {
		ServiceName       string `yaml:"service-name"`
		DockerImageName   string `yaml:"docker-image-name"`
		ScreenSessionName string `yaml:"screen-session-name"`
	}

	SuimonConfig struct {
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
			SystemTable struct {
				Display bool `yaml:"display"`
			} `yaml:"system-table"`
			ValidatorsTable struct {
				Display bool `yaml:"display"`
			} `yaml:"validators-table"`
		} `yaml:"monitors-config"`
		PublicRPC struct {
			Testnet []string `yaml:"testnet"`
			Devnet  []string `yaml:"devnet"`
		} `yaml:"public-rpc"`
		NodeConfigPath string `yaml:"node-config-path"`
		Network        struct {
			Name        string `yaml:"name"`
			NetworkType enums.NetworkType
		} `yaml:"network"`
		IPLookup struct {
			AccessToken string `yaml:"access-token"`
		} `yaml:"ip-lookup"`
		MonitorsVisual struct {
			ColorScheme  string `yaml:"color-scheme"`
			EnableEmojis bool   `yaml:"enable-emojis"`
		} `yaml:"monitors-visual"`
		ProcessLaunchType ProcessLaunchType `yaml:"process-launch-type"`
	}
)

// ParseSuimonConfig decodes the Suimon configuration file at the given path and returns a pointer to a populated SuimonConfig struct.
// This function accepts the following parameter:
// - path: a pointer to a string representing the file path of the Suimon configuration file to be parsed.
// The function returns a pointer to a SuimonConfig struct containing the parsed configuration data, and an error if there was an issue parsing the configuration file.
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

	result.SetProcessLaunchType()

	return &result, nil
}

// SetProcessLaunchType sets the launch type for the SuiNode process based on the value specified in the configuration.
// The launch type can be specified as one of the following: "service", "docker", or "screen".
func (sconfig *SuimonConfig) SetProcessLaunchType() {
	processLaunchType := sconfig.ProcessLaunchType
	if processLaunchType.ServiceName == "" && processLaunchType.DockerImageName == "" && processLaunchType.ScreenSessionName == "" {
		sconfig.ProcessLaunchType = ProcessLaunchType{
			ServiceName:       defaultServiceName,
			DockerImageName:   defaultImageName,
			ScreenSessionName: defaultScreenName,
		}
	}
}

// SetNetworkConfig sets the network configuration for the SuimonConfig struct to the given network type.
// This function accepts the following parameter:
// - networkType: an enums.NetworkType representing the type of network to configure the SuimonConfig struct for.
// This function does not return anything.
func (sconfig *SuimonConfig) SetNetworkConfig(networkType enums.NetworkType) {
	sconfig.Network.NetworkType = networkType
	sconfig.Network.Name = networkType.String()
}

// GetRPCByNetwork returns a list of RPC endpoint strings for the network configured in the SuimonConfig struct.
// This function does not accept any parameters.
// The function returns a slice of strings representing the RPC endpoints for the configured network.
func (sconfig *SuimonConfig) GetRPCByNetwork() []string {
	networkType := sconfig.Network.NetworkType

	switch networkType {
	case enums.NetworkTypeDevnet:
		return sconfig.PublicRPC.Devnet
	case enums.NetworkTypeTestnet:
		return sconfig.PublicRPC.Testnet
	}

	return nil
}
