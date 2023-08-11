package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgGetStorage = "get_storage"

var _ sdk.Msg = &MsgGetStorage{}

func NewMsgGetStorage(creator string, address string, position string, blockTag string) *MsgGetStorage {
	return &MsgGetStorage{
		Creator:  creator,
		Address:  address,
		Position: position,
		BlockTag: blockTag,
	}
}

func (msg *MsgGetStorage) Route() string {
	return RouterKey
}

func (msg *MsgGetStorage) Type() string {
	return TypeMsgGetStorage
}

func (msg *MsgGetStorage) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgGetStorage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgGetStorage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
