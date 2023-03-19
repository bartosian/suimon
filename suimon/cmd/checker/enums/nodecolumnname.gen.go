// Code generated by "enumer -type=NodeColumnName -json -transform=snake-upper -output=./nodecolumnname.gen.go"; DO NOT EDIT.

package enums

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _NodeColumnNameName = "COLUMN_NAME_STATUSCOLUMN_NAME_ADDRESSCOLUMN_NAME_PORT_RPCCOLUMN_NAME_TOTAL_TRANSACTIONSCOLUMN_NAME_LATEST_CHECKPOINTCOLUMN_NAME_HIGHEST_CHECKPOINTCOLUMN_NAME_TX_SYNC_PROGRESSCOLUMN_NAME_CHECK_SYNC_PROGRESSCOLUMN_NAME_CONNECTED_PEERSCOLUMN_NAME_UPTIMECOLUMN_NAME_VERSIONCOLUMN_NAME_COMMITCOLUMN_NAME_COMPANYCOLUMN_NAME_COUNTRY"

var _NodeColumnNameIndex = [...]uint16{0, 18, 37, 57, 87, 116, 146, 174, 205, 232, 250, 269, 287, 306, 325}

const _NodeColumnNameLowerName = "column_name_statuscolumn_name_addresscolumn_name_port_rpccolumn_name_total_transactionscolumn_name_latest_checkpointcolumn_name_highest_checkpointcolumn_name_tx_sync_progresscolumn_name_check_sync_progresscolumn_name_connected_peerscolumn_name_uptimecolumn_name_versioncolumn_name_commitcolumn_name_companycolumn_name_country"

func (i NodeColumnName) String() string {
	if i < 0 || i >= NodeColumnName(len(_NodeColumnNameIndex)-1) {
		return fmt.Sprintf("NodeColumnName(%d)", i)
	}
	return _NodeColumnNameName[_NodeColumnNameIndex[i]:_NodeColumnNameIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _NodeColumnNameNoOp() {
	var x [1]struct{}
	_ = x[ColumnNameStatus-(0)]
	_ = x[ColumnNameAddress-(1)]
	_ = x[ColumnNamePortRPC-(2)]
	_ = x[ColumnNameTotalTransactions-(3)]
	_ = x[ColumnNameLatestCheckpoint-(4)]
	_ = x[ColumnNameHighestCheckpoint-(5)]
	_ = x[ColumnNameTXSyncProgress-(6)]
	_ = x[ColumnNameCheckSyncProgress-(7)]
	_ = x[ColumnNameConnectedPeers-(8)]
	_ = x[ColumnNameUptime-(9)]
	_ = x[ColumnNameVersion-(10)]
	_ = x[ColumnNameCommit-(11)]
	_ = x[ColumnNameCompany-(12)]
	_ = x[ColumnNameCountry-(13)]
}

var _NodeColumnNameValues = []NodeColumnName{ColumnNameStatus, ColumnNameAddress, ColumnNamePortRPC, ColumnNameTotalTransactions, ColumnNameLatestCheckpoint, ColumnNameHighestCheckpoint, ColumnNameTXSyncProgress, ColumnNameCheckSyncProgress, ColumnNameConnectedPeers, ColumnNameUptime, ColumnNameVersion, ColumnNameCommit, ColumnNameCompany, ColumnNameCountry}

var _NodeColumnNameNameToValueMap = map[string]NodeColumnName{
	_NodeColumnNameName[0:18]:         ColumnNameStatus,
	_NodeColumnNameLowerName[0:18]:    ColumnNameStatus,
	_NodeColumnNameName[18:37]:        ColumnNameAddress,
	_NodeColumnNameLowerName[18:37]:   ColumnNameAddress,
	_NodeColumnNameName[37:57]:        ColumnNamePortRPC,
	_NodeColumnNameLowerName[37:57]:   ColumnNamePortRPC,
	_NodeColumnNameName[57:87]:        ColumnNameTotalTransactions,
	_NodeColumnNameLowerName[57:87]:   ColumnNameTotalTransactions,
	_NodeColumnNameName[87:116]:       ColumnNameLatestCheckpoint,
	_NodeColumnNameLowerName[87:116]:  ColumnNameLatestCheckpoint,
	_NodeColumnNameName[116:146]:      ColumnNameHighestCheckpoint,
	_NodeColumnNameLowerName[116:146]: ColumnNameHighestCheckpoint,
	_NodeColumnNameName[146:174]:      ColumnNameTXSyncProgress,
	_NodeColumnNameLowerName[146:174]: ColumnNameTXSyncProgress,
	_NodeColumnNameName[174:205]:      ColumnNameCheckSyncProgress,
	_NodeColumnNameLowerName[174:205]: ColumnNameCheckSyncProgress,
	_NodeColumnNameName[205:232]:      ColumnNameConnectedPeers,
	_NodeColumnNameLowerName[205:232]: ColumnNameConnectedPeers,
	_NodeColumnNameName[232:250]:      ColumnNameUptime,
	_NodeColumnNameLowerName[232:250]: ColumnNameUptime,
	_NodeColumnNameName[250:269]:      ColumnNameVersion,
	_NodeColumnNameLowerName[250:269]: ColumnNameVersion,
	_NodeColumnNameName[269:287]:      ColumnNameCommit,
	_NodeColumnNameLowerName[269:287]: ColumnNameCommit,
	_NodeColumnNameName[287:306]:      ColumnNameCompany,
	_NodeColumnNameLowerName[287:306]: ColumnNameCompany,
	_NodeColumnNameName[306:325]:      ColumnNameCountry,
	_NodeColumnNameLowerName[306:325]: ColumnNameCountry,
}

var _NodeColumnNameNames = []string{
	_NodeColumnNameName[0:18],
	_NodeColumnNameName[18:37],
	_NodeColumnNameName[37:57],
	_NodeColumnNameName[57:87],
	_NodeColumnNameName[87:116],
	_NodeColumnNameName[116:146],
	_NodeColumnNameName[146:174],
	_NodeColumnNameName[174:205],
	_NodeColumnNameName[205:232],
	_NodeColumnNameName[232:250],
	_NodeColumnNameName[250:269],
	_NodeColumnNameName[269:287],
	_NodeColumnNameName[287:306],
	_NodeColumnNameName[306:325],
}

// NodeColumnNameString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func NodeColumnNameString(s string) (NodeColumnName, error) {
	if val, ok := _NodeColumnNameNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _NodeColumnNameNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to NodeColumnName values", s)
}

// NodeColumnNameValues returns all values of the enum
func NodeColumnNameValues() []NodeColumnName {
	return _NodeColumnNameValues
}

// NodeColumnNameStrings returns a slice of all String values of the enum
func NodeColumnNameStrings() []string {
	strs := make([]string, len(_NodeColumnNameNames))
	copy(strs, _NodeColumnNameNames)
	return strs
}

// IsANodeColumnName returns "true" if the value is listed in the enum definition. "false" otherwise
func (i NodeColumnName) IsANodeColumnName() bool {
	for _, v := range _NodeColumnNameValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for NodeColumnName
func (i NodeColumnName) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for NodeColumnName
func (i *NodeColumnName) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("NodeColumnName should be a string, got %s", data)
	}

	var err error
	*i, err = NodeColumnNameString(s)
	return err
}