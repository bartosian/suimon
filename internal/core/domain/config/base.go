package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strings"
)

const (
	suimonConfigEnvVar = "SUIMON_CONFIG_PATH"
	suimonConfigDir    = ".suimon"
	ymlPattern         = "suimon-*.yml"
	yamlPattern        = "suimon-*.yaml"
)

type Config struct {
	PublicRPC []string `yaml:"public-rpc"`
	FullNodes []struct {
		JSONRPCAddress string `yaml:"json-rpc-address"`
		MetricsAddress string `yaml:"metrics-address"`
	} `yaml:"full-nodes"`
	Validators []struct {
		MetricsAddress string `yaml:"metrics-address"`
	} `yaml:"validators"`
	IPLookup struct {
		AccessToken string `yaml:"access-token"`
	} `yaml:"ip-lookup"`
}

// NewConfig reads the Suimon configuration files from the directory specified by
// the SUIMON_CONFIG_PATH environment variable or the default directory if the
// environment variable is not set, and returns a map of Config objects with the
// file name segments as the keys.
func NewConfig() (map[string]Config, error) {
	dirPath := os.Getenv(suimonConfigEnvVar)
	if dirPath == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}

		dirPath = filepath.Join(homeDir, suimonConfigDir)
	}

	return readConfigs(dirPath)
}

// readConfigs reads the Suimon configuration files from the specified directory,
// creates a map of Config objects with the file name segments as the keys, and returns
// the map. The file name segments are converted to uppercase before being used as keys.
func readConfigs(dirPath string) (map[string]Config, error) {
	configs := make(map[string]Config)

	ymlFiles, _ := filepath.Glob(filepath.Join(dirPath, ymlPattern))
	yamlFiles, _ := filepath.Glob(filepath.Join(dirPath, yamlPattern))

	if len(ymlFiles)+len(yamlFiles) == 0 {
		return nil, fmt.Errorf("no suimon configuration files found in %s", dirPath)
	}

	for _, file := range append(ymlFiles, yamlFiles...) {
		fileData, err := os.ReadFile(file)
		if err != nil {
			return nil, err
		}

		var config Config
		err = yaml.Unmarshal(fileData, &config)
		if err != nil {
			return nil, err
		}

		filename := filepath.Base(file)
		filename = strings.TrimPrefix(filename, "suimon-")
		filename = strings.TrimSuffix(filename, ".yml")
		filename = strings.TrimSuffix(filename, ".yaml")
		filename = strings.ToUpper(filename)

		configs[filename] = config
	}

	return configs, nil
}
