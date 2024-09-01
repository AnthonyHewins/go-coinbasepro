// Code generated by "enumer -type PortfolioType -transform snake_upper -trimprefix PortfolioType"; DO NOT EDIT.

package coinbase

import (
	"fmt"
	"strings"
)

const _PortfolioTypeName = "UNDEFINEDDEFAULTCONSUMERINTX"

var _PortfolioTypeIndex = [...]uint8{0, 9, 16, 24, 28}

const _PortfolioTypeLowerName = "undefineddefaultconsumerintx"

func (i PortfolioType) String() string {
	if i >= PortfolioType(len(_PortfolioTypeIndex)-1) {
		return fmt.Sprintf("PortfolioType(%d)", i)
	}
	return _PortfolioTypeName[_PortfolioTypeIndex[i]:_PortfolioTypeIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _PortfolioTypeNoOp() {
	var x [1]struct{}
	_ = x[PortfolioTypeUndefined-(0)]
	_ = x[PortfolioTypeDefault-(1)]
	_ = x[PortfolioTypeConsumer-(2)]
	_ = x[PortfolioTypeIntx-(3)]
}

var _PortfolioTypeValues = []PortfolioType{PortfolioTypeUndefined, PortfolioTypeDefault, PortfolioTypeConsumer, PortfolioTypeIntx}

var _PortfolioTypeNameToValueMap = map[string]PortfolioType{
	_PortfolioTypeName[0:9]:        PortfolioTypeUndefined,
	_PortfolioTypeLowerName[0:9]:   PortfolioTypeUndefined,
	_PortfolioTypeName[9:16]:       PortfolioTypeDefault,
	_PortfolioTypeLowerName[9:16]:  PortfolioTypeDefault,
	_PortfolioTypeName[16:24]:      PortfolioTypeConsumer,
	_PortfolioTypeLowerName[16:24]: PortfolioTypeConsumer,
	_PortfolioTypeName[24:28]:      PortfolioTypeIntx,
	_PortfolioTypeLowerName[24:28]: PortfolioTypeIntx,
}

var _PortfolioTypeNames = []string{
	_PortfolioTypeName[0:9],
	_PortfolioTypeName[9:16],
	_PortfolioTypeName[16:24],
	_PortfolioTypeName[24:28],
}

// PortfolioTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func PortfolioTypeString(s string) (PortfolioType, error) {
	if val, ok := _PortfolioTypeNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _PortfolioTypeNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to PortfolioType values", s)
}

// PortfolioTypeValues returns all values of the enum
func PortfolioTypeValues() []PortfolioType {
	return _PortfolioTypeValues
}

// PortfolioTypeStrings returns a slice of all String values of the enum
func PortfolioTypeStrings() []string {
	strs := make([]string, len(_PortfolioTypeNames))
	copy(strs, _PortfolioTypeNames)
	return strs
}

// IsAPortfolioType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i PortfolioType) IsAPortfolioType() bool {
	for _, v := range _PortfolioTypeValues {
		if i == v {
			return true
		}
	}
	return false
}
