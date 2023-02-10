package enums

import (
	"fmt"
	"strings"
)

const (
	NetworkTypeUndefined NetworkType = "UNDEFINED"
	NetworkTypeDevnet    NetworkType = "DEVNET"
	NetworkTypeTestnet   NetworkType = "TESTNET"
)

type NetworkType string

const (
	devnetRPC  = "https://fullnode.devnet.sui.io:443"
	testnetRPC = "https://fullnode.testnet.sui.io:443"
)

var networkTypeValues = [...]string{
	"DEVNET",
	"TESTNET",
}

func (i NetworkType) String() string {
	return string(i)
}

func NetworkTypeFromString(value string) (NetworkType, error) {
	value = strings.ToUpper(strings.TrimSpace(value))

	result, ok := map[string]NetworkType{
		"DEVNET":  NetworkTypeDevnet,
		"TESTNET": NetworkTypeTestnet,
	}[value]

	if ok {
		return result, nil
	}

	return NetworkTypeUndefined, fmt.Errorf("unsupported network type enum string: %s", value)
}

func (i NetworkType) ToRPC() string {
	switch i {
	case NetworkTypeTestnet:
		return testnetRPC
	default:
		return devnetRPC
	}
}
