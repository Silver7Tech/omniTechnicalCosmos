package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		OmniList: []Omni{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in omni
	omniIdMap := make(map[uint64]bool)
	omniCount := gs.GetOmniCount()
	for _, elem := range gs.OmniList {
		if _, ok := omniIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for omni")
		}
		if elem.Id >= omniCount {
			return fmt.Errorf("omni id should be lower or equal than the last id")
		}
		omniIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
