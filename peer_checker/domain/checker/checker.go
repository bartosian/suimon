package checker

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/ybbus/jsonrpc/v3"
	"gopkg.in/yaml.v3"
)

const (
	peerSeparator = "/"
	peerCount     = 4
)

type (
	PeerData struct {
		Address string `yaml:"address"`
	}

	P2PConfig struct {
		SeedPeers []PeerData `yaml:"seed-peers"`
	}

	NodeConfigYaml struct {
		P2PConfig P2PConfig `yaml:"p2p-config"`
	}

	Checker struct {
		Peers     []Peer
		rpcClient jsonrpc.RPCClient
	}
)

func NewChecker(path string) (*Checker, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var configData NodeConfigYaml

	err = yaml.Unmarshal(file, &configData)
	if err != nil {
		return nil, err
	}

	peers, err := configData.P2PConfig.MapToCheckerPeers()
	if err != nil {
		return nil, err
	}

	return &Checker{
		Peers: peers,
	}, nil
}

func (config *P2PConfig) MapToCheckerPeers() ([]Peer, error) {
	configPeers := config.SeedPeers
	checkerPeers := make([]Peer, 0, len(configPeers))

	for _, peer := range configPeers {
		if isValidCharCount(peer.Address, peerSeparator, peerCount) {
			peerInfo := strings.Split(peer.Address, peerSeparator)

			checkerPeer, err := newPeer(peerInfo[2], peerInfo[4])
			if err != nil {
				return nil, err
			}

			checkerPeers = append(checkerPeers, *checkerPeer)

			continue
		}

		return nil, fmt.Errorf("invalid peer address value provided %s", peer.Address)
	}

	if len(checkerPeers) == 0 {
		return nil, errors.New("no peers found in config file")
	}

	return checkerPeers, nil
}
