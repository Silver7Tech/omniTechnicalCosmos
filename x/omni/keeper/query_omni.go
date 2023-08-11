package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"omni/x/omni/types"
)

func (k Keeper) OmniAll(goCtx context.Context, req *types.QueryAllOmniRequest) (*types.QueryAllOmniResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var omnis []types.Omni
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	omniStore := prefix.NewStore(store, types.KeyPrefix(types.OmniKey))

	pageRes, err := query.Paginate(omniStore, req.Pagination, func(key []byte, value []byte) error {
		var omni types.Omni
		if err := k.cdc.Unmarshal(value, &omni); err != nil {
			return err
		}

		omnis = append(omnis, omni)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllOmniResponse{Omni: omnis, Pagination: pageRes}, nil
}

func (k Keeper) Omni(goCtx context.Context, req *types.QueryGetOmniRequest) (*types.QueryGetOmniResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	omni, found := k.GetOmni(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetOmniResponse{Omni: omni}, nil
}
