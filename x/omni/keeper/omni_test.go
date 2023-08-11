package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "omni/testutil/keeper"
	"omni/testutil/nullify"
	"omni/x/omni/keeper"
	"omni/x/omni/types"
)

func createNOmni(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Omni {
	items := make([]types.Omni, n)
	for i := range items {
		items[i].Id = keeper.AppendOmni(ctx, items[i])
	}
	return items
}

func TestOmniGet(t *testing.T) {
	keeper, ctx := keepertest.OmniKeeper(t)
	items := createNOmni(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetOmni(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestOmniRemove(t *testing.T) {
	keeper, ctx := keepertest.OmniKeeper(t)
	items := createNOmni(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveOmni(ctx, item.Id)
		_, found := keeper.GetOmni(ctx, item.Id)
		require.False(t, found)
	}
}

func TestOmniGetAll(t *testing.T) {
	keeper, ctx := keepertest.OmniKeeper(t)
	items := createNOmni(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllOmni(ctx)),
	)
}

func TestOmniCount(t *testing.T) {
	keeper, ctx := keepertest.OmniKeeper(t)
	items := createNOmni(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetOmniCount(ctx))
}
