package omni_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "omni/testutil/keeper"
	"omni/testutil/nullify"
	"omni/x/omni"
	"omni/x/omni/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		OmniList: []types.Omni{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		OmniCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OmniKeeper(t)
	omni.InitGenesis(ctx, *k, genesisState)
	got := omni.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.OmniList, got.OmniList)
	require.Equal(t, genesisState.OmniCount, got.OmniCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
