package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PythonCodeList: []PythonCode{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in pythonCode
	pythonCodeIdMap := make(map[uint64]bool)
	pythonCodeCount := gs.GetPythonCodeCount()
	for _, elem := range gs.PythonCodeList {
		if _, ok := pythonCodeIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for pythonCode")
		}
		if elem.Id >= pythonCodeCount {
			return fmt.Errorf("pythonCode id should be lower or equal than the last id")
		}
		pythonCodeIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
