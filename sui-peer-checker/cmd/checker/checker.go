package checker

import (
	"os"

	"github.com/ybbus/jsonrpc/v3"
	"gopkg.in/yaml.v3"

	"github.com/bartosian/sui_helpers/sui-peer-checker/cmd/checker/enums"
	"github.com/bartosian/sui_helpers/sui-peer-checker/cmd/checker/tablebuilder"
	"github.com/bartosian/sui_helpers/sui-peer-checker/cmd/checker/tablebuilder/tables"
)

const (
	peerSeparator = "/"
	peerCount     = 4
)

type (
	PeerData struct {
		Address string `yaml:"address"`
	}

	NodeYaml struct {
		Config Config `yaml:"p2p-config"`
	}

	Checker struct {
		peers        []Peer
		rpcClient    jsonrpc.RPCClient
		tableBuilder *tablebuilder.TableBuilder
		tableConfig  tablebuilder.TableConfig
	}
)

func NewChecker(path string, network enums.NetworkType) (*Checker, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var nodeYaml NodeYaml
	err = yaml.Unmarshal(file, &nodeYaml)
	if err != nil {
		return nil, err
	}

	peers, err := nodeYaml.Config.parsePeers()
	if err != nil {
		return nil, err
	}

	return &Checker{
		peers:     peers,
		rpcClient: jsonrpc.NewClient(network.ToRPC()),
	}, nil
}

func (checker *Checker) GenerateTableConfig() {
	tableConfig := tablebuilder.TableConfig{
		Name:         tables.TableTitleSUI,
		Style:        tables.TableStyleSUI,
		RowsCount:    len(checker.peers),
		ColumnsCount: len(tables.ColumnConfigSUI),
		SortConfig:   tables.TableSortConfigSUI,
	}

	columns := make([]tablebuilder.Column, len(tables.ColumnConfigSUI))

	for idx, config := range tables.ColumnConfigSUI {
		columns[idx].Config = config
	}

	for _, peer := range checker.peers {
		columns[tables.ColumnNameSUIPeer].SetValue(peer.Address)
		columns[tables.ColumnNameSUIPort].SetValue(peer.Port)
		columns[tables.ColumnNameSUITotalTransactions].SetValue(peer.Metrics.TotalTransactionNumber)
		columns[tables.ColumnNameSUIHighestCheckpoints].SetValue(peer.Metrics.HighestSyncedCheckpoint)
		columns[tables.ColumnNameSUIConnectedPeers].SetValue(peer.Metrics.SuiNetworkPeers)
		columns[tables.ColumnNameSUIUptime].SetValue(peer.Metrics.Uptime)
		columns[tables.ColumnNameSUIVersion].SetValue(peer.Metrics.Version)
		columns[tables.ColumnNameSUICommit].SetValue(peer.Metrics.Commit)
		columns[tables.ColumnNameSUICountry].SetValue(peer.Location.String())
	}

	tableConfig.Columns = columns

	checker.tableBuilder = tablebuilder.NewTableBuilder(tableConfig)
}
