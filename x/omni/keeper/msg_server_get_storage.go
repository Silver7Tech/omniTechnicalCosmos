package keeper

import (
	"context"

	"omni/x/ethservice"
	"omni/x/omni/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) GetStorage(goCtx context.Context, msg *types.MsgGetStorage) (*types.MsgGetStorageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	ethService, err := ethservice.NewEthService("https://mainnet.infura.io/v3/46d7e5b569734fad998bf16a3c13b3e4")
	if err != nil {
		return nil, err
	}
	storage, err := ethService.GetStorageAt(msg.Address, msg.Position, msg.BlockTag)
	var omni = types.Omni{
		Address:  msg.Address,
		Position: msg.Position,
		BlockTag: msg.BlockTag,
		Storage:  storage,
	}
	k.AppendOmni(ctx, omni)
	return &types.MsgGetStorageResponse{}, nil
}
