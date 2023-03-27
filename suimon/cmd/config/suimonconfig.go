package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

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
			ValidatorTable struct {
				Display bool `yaml:"display"`
			} `yaml:"validator-table"`
			PeersTable struct {
				Display bool `yaml:"display"`
			} `yaml:"peers-table"`
			SystemTable struct {
				Display bool `yaml:"display"`
			} `yaml:"system-table"`
			ActiveValidatorsTable struct {
				Display bool `yaml:"display"`
			} `yaml:"active-validators-table"`
		} `yaml:"monitors-config"`
		PublicRPC []string `yaml:"public-rpc"`
		FullNode  struct {
			JSONRPCAddress string `yaml:"json-rpc-address"`
			MetricsAddress string `yaml:"metrics-address"`
		} `yaml:"full-node"`
		Validator struct {
			MetricsAddress string `yaml:"metrics-address"`
		} `yaml:"validator"`
		SeedPeers []string `yaml:"seed-peers"`
		IPLookup  struct {
			AccessToken string `yaml:"access-token"`
		} `yaml:"ip-lookup"`
		MonitorsVisual struct {
			EnableEmojis bool `yaml:"enable-emojis"`
		} `yaml:"monitors-visual"`
		DbPath            string            `yaml:"db-path"`
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
