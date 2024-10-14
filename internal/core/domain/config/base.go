package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	suimonConfigEnvVar = "SUIMON_CONFIG_PATH"
	suimonConfigDir    = ".suimon"
)

type Config struct {
	IPLookup struct {
		AccessToken string `yaml:"access-token"`
	} `yaml:"ip-lookup"`
	PublicRPC []string `yaml:"public-rpc"`
	FullNodes []struct {
		JSONRPCAddress string `yaml:"json-rpc-address"`
		MetricsAddress string `yaml:"metrics-address"`
	} `yaml:"full-nodes"`
	Validators []struct {
		MetricsAddress string `yaml:"metrics-address"`
	} `yaml:"validators"`
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

	// Retrieve .yml files
	ymlFiles, err := filepath.Glob(filepath.Join(dirPath, "*.yml"))
	if err != nil {
		return nil, err
	}

	// Retrieve .yaml files
	yamlFiles, err := filepath.Glob(filepath.Join(dirPath, "*.yaml"))
	if err != nil {
		return nil, err
	}

	// Combine file lists
	ymlFiles = append(ymlFiles, yamlFiles...)

	if len(ymlFiles) == 0 {
		return nil, fmt.Errorf("no Suimon configuration files found in %s", dirPath)
	}

	for _, file := range ymlFiles {
		fileData, readErr := os.ReadFile(file)
		if readErr != nil {
			return nil, fmt.Errorf("error reading file %s: %w", file, readErr)
		}

		var config Config
		if unmarshalErr := yaml.Unmarshal(fileData, &config); unmarshalErr != nil {
			return nil, fmt.Errorf("error unmarshaling YAML in file %s: %w", file, unmarshalErr)
		}

		filename := strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
		filename = strings.TrimPrefix(filename, "suimon-")
		filename = strings.ToUpper(filename)

		configs[filename] = config
	}

	return configs, nil
}
