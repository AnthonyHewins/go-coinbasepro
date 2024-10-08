// Code generated by "enumer -type FuturesPosition -json -transform snake_upper"; DO NOT EDIT.

package coinbase

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _FuturesPositionName = "FUTURES_POSITION_SIDE_UNSPECIFIEDFUTURES_POSITION_SIDE_LONGFUTURES_POSITION_SIDE_SHORT"

var _FuturesPositionIndex = [...]uint8{0, 33, 59, 86}

const _FuturesPositionLowerName = "futures_position_side_unspecifiedfutures_position_side_longfutures_position_side_short"

func (i FuturesPosition) String() string {
	if i >= FuturesPosition(len(_FuturesPositionIndex)-1) {
		return fmt.Sprintf("FuturesPosition(%d)", i)
	}
	return _FuturesPositionName[_FuturesPositionIndex[i]:_FuturesPositionIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _FuturesPositionNoOp() {
	var x [1]struct{}
	_ = x[FuturesPositionSideUnspecified-(0)]
	_ = x[FuturesPositionSideLong-(1)]
	_ = x[FuturesPositionSideShort-(2)]
}

var _FuturesPositionValues = []FuturesPosition{FuturesPositionSideUnspecified, FuturesPositionSideLong, FuturesPositionSideShort}

var _FuturesPositionNameToValueMap = map[string]FuturesPosition{
	_FuturesPositionName[0:33]:       FuturesPositionSideUnspecified,
	_FuturesPositionLowerName[0:33]:  FuturesPositionSideUnspecified,
	_FuturesPositionName[33:59]:      FuturesPositionSideLong,
	_FuturesPositionLowerName[33:59]: FuturesPositionSideLong,
	_FuturesPositionName[59:86]:      FuturesPositionSideShort,
	_FuturesPositionLowerName[59:86]: FuturesPositionSideShort,
}

var _FuturesPositionNames = []string{
	_FuturesPositionName[0:33],
	_FuturesPositionName[33:59],
	_FuturesPositionName[59:86],
}

// FuturesPositionString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func FuturesPositionString(s string) (FuturesPosition, error) {
	if val, ok := _FuturesPositionNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _FuturesPositionNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to FuturesPosition values", s)
}

// FuturesPositionValues returns all values of the enum
func FuturesPositionValues() []FuturesPosition {
	return _FuturesPositionValues
}

// FuturesPositionStrings returns a slice of all String values of the enum
func FuturesPositionStrings() []string {
	strs := make([]string, len(_FuturesPositionNames))
	copy(strs, _FuturesPositionNames)
	return strs
}

// IsAFuturesPosition returns "true" if the value is listed in the enum definition. "false" otherwise
func (i FuturesPosition) IsAFuturesPosition() bool {
	for _, v := range _FuturesPositionValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for FuturesPosition
func (i FuturesPosition) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for FuturesPosition
func (i *FuturesPosition) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("FuturesPosition should be a string, got %s", data)
	}

	var err error
	*i, err = FuturesPositionString(s)
	return err
}
