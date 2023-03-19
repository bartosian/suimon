// Code generated by "enumer -type=ValidatorColumnName -json -transform=snake-upper -output=./validatorcolumnname.gen.go"; DO NOT EDIT.

package enums

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _ValidatorColumnNameName = "VALIDATOR_COLUMN_NAME_NAMEVALIDATOR_COLUMN_NAME_ADDRESSVALIDATOR_COLUMN_NAME_VOTING_POWERVALIDATOR_COLUMN_NAME_GAS_PRICEVALIDATOR_COLUMN_NAME_STAKING_POOL_SUI_BALANCEVALIDATOR_COLUMN_NAME_COMMISSION_RATEVALIDATOR_COLUMN_NAME_NEXT_EPOCH_GAS_PRICEVALIDATOR_COLUMN_NAME_NEXT_EPOCH_STAKEVALIDATOR_COLUMN_NAME_NEXT_EPOCH_COMMISSION_RATE"

var _ValidatorColumnNameIndex = [...]uint16{0, 26, 55, 89, 120, 166, 203, 245, 283, 331}

const _ValidatorColumnNameLowerName = "validator_column_name_namevalidator_column_name_addressvalidator_column_name_voting_powervalidator_column_name_gas_pricevalidator_column_name_staking_pool_sui_balancevalidator_column_name_commission_ratevalidator_column_name_next_epoch_gas_pricevalidator_column_name_next_epoch_stakevalidator_column_name_next_epoch_commission_rate"

func (i ValidatorColumnName) String() string {
	if i < 0 || i >= ValidatorColumnName(len(_ValidatorColumnNameIndex)-1) {
		return fmt.Sprintf("ValidatorColumnName(%d)", i)
	}
	return _ValidatorColumnNameName[_ValidatorColumnNameIndex[i]:_ValidatorColumnNameIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _ValidatorColumnNameNoOp() {
	var x [1]struct{}
	_ = x[ValidatorColumnNameName-(0)]
	_ = x[ValidatorColumnNameAddress-(1)]
	_ = x[ValidatorColumnNameVotingPower-(2)]
	_ = x[ValidatorColumnNameGasPrice-(3)]
	_ = x[ValidatorColumnNameStakingPoolSuiBalance-(4)]
	_ = x[ValidatorColumnNameCommissionRate-(5)]
	_ = x[ValidatorColumnNameNextEpochGasPrice-(6)]
	_ = x[ValidatorColumnNameNextEpochStake-(7)]
	_ = x[ValidatorColumnNameNextEpochCommissionRate-(8)]
}

var _ValidatorColumnNameValues = []ValidatorColumnName{ValidatorColumnNameName, ValidatorColumnNameAddress, ValidatorColumnNameVotingPower, ValidatorColumnNameGasPrice, ValidatorColumnNameStakingPoolSuiBalance, ValidatorColumnNameCommissionRate, ValidatorColumnNameNextEpochGasPrice, ValidatorColumnNameNextEpochStake, ValidatorColumnNameNextEpochCommissionRate}

var _ValidatorColumnNameNameToValueMap = map[string]ValidatorColumnName{
	_ValidatorColumnNameName[0:26]:         ValidatorColumnNameName,
	_ValidatorColumnNameLowerName[0:26]:    ValidatorColumnNameName,
	_ValidatorColumnNameName[26:55]:        ValidatorColumnNameAddress,
	_ValidatorColumnNameLowerName[26:55]:   ValidatorColumnNameAddress,
	_ValidatorColumnNameName[55:89]:        ValidatorColumnNameVotingPower,
	_ValidatorColumnNameLowerName[55:89]:   ValidatorColumnNameVotingPower,
	_ValidatorColumnNameName[89:120]:       ValidatorColumnNameGasPrice,
	_ValidatorColumnNameLowerName[89:120]:  ValidatorColumnNameGasPrice,
	_ValidatorColumnNameName[120:166]:      ValidatorColumnNameStakingPoolSuiBalance,
	_ValidatorColumnNameLowerName[120:166]: ValidatorColumnNameStakingPoolSuiBalance,
	_ValidatorColumnNameName[166:203]:      ValidatorColumnNameCommissionRate,
	_ValidatorColumnNameLowerName[166:203]: ValidatorColumnNameCommissionRate,
	_ValidatorColumnNameName[203:245]:      ValidatorColumnNameNextEpochGasPrice,
	_ValidatorColumnNameLowerName[203:245]: ValidatorColumnNameNextEpochGasPrice,
	_ValidatorColumnNameName[245:283]:      ValidatorColumnNameNextEpochStake,
	_ValidatorColumnNameLowerName[245:283]: ValidatorColumnNameNextEpochStake,
	_ValidatorColumnNameName[283:331]:      ValidatorColumnNameNextEpochCommissionRate,
	_ValidatorColumnNameLowerName[283:331]: ValidatorColumnNameNextEpochCommissionRate,
}

var _ValidatorColumnNameNames = []string{
	_ValidatorColumnNameName[0:26],
	_ValidatorColumnNameName[26:55],
	_ValidatorColumnNameName[55:89],
	_ValidatorColumnNameName[89:120],
	_ValidatorColumnNameName[120:166],
	_ValidatorColumnNameName[166:203],
	_ValidatorColumnNameName[203:245],
	_ValidatorColumnNameName[245:283],
	_ValidatorColumnNameName[283:331],
}

// ValidatorColumnNameString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ValidatorColumnNameString(s string) (ValidatorColumnName, error) {
	if val, ok := _ValidatorColumnNameNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _ValidatorColumnNameNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to ValidatorColumnName values", s)
}

// ValidatorColumnNameValues returns all values of the enum
func ValidatorColumnNameValues() []ValidatorColumnName {
	return _ValidatorColumnNameValues
}

// ValidatorColumnNameStrings returns a slice of all String values of the enum
func ValidatorColumnNameStrings() []string {
	strs := make([]string, len(_ValidatorColumnNameNames))
	copy(strs, _ValidatorColumnNameNames)
	return strs
}

// IsAValidatorColumnName returns "true" if the value is listed in the enum definition. "false" otherwise
func (i ValidatorColumnName) IsAValidatorColumnName() bool {
	for _, v := range _ValidatorColumnNameValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for ValidatorColumnName
func (i ValidatorColumnName) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for ValidatorColumnName
func (i *ValidatorColumnName) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("ValidatorColumnName should be a string, got %s", data)
	}

	var err error
	*i, err = ValidatorColumnNameString(s)
	return err
}