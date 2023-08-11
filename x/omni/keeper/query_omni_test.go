package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "omni/testutil/keeper"
	"omni/testutil/nullify"
	"omni/x/omni/types"
)

func TestOmniQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.OmniKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNOmni(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetOmniRequest
		response *types.QueryGetOmniResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetOmniRequest{Id: msgs[0].Id},
			response: &types.QueryGetOmniResponse{Omni: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetOmniRequest{Id: msgs[1].Id},
			response: &types.QueryGetOmniResponse{Omni: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetOmniRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Omni(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestOmniQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.OmniKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNOmni(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllOmniRequest {
		return &types.QueryAllOmniRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.OmniAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Omni), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Omni),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.OmniAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Omni), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Omni),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.OmniAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Omni),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.OmniAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
