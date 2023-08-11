package omni

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"omni/x/omni/keeper"
	"omni/x/omni/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the omni
	for _, elem := range genState.OmniList {
		k.SetOmni(ctx, elem)
	}

	// Set omni count
	k.SetOmniCount(ctx, genState.OmniCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.OmniList = k.GetAllOmni(ctx)
	genesis.OmniCount = k.GetOmniCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
