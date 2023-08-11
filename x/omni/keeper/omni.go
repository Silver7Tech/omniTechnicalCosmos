package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"omni/x/omni/types"
)

// GetOmniCount get the total number of omni
func (k Keeper) GetOmniCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.OmniCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetOmniCount set the total number of omni
func (k Keeper) SetOmniCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.OmniCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendOmni appends a omni in the store with a new id and update the count
func (k Keeper) AppendOmni(
	ctx sdk.Context,
	omni types.Omni,
) uint64 {
	// Create the omni
	count := k.GetOmniCount(ctx)

	// Set the ID of the appended value
	omni.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OmniKey))
	appendedValue := k.cdc.MustMarshal(&omni)
	store.Set(GetOmniIDBytes(omni.Id), appendedValue)

	// Update omni count
	k.SetOmniCount(ctx, count+1)

	return count
}

// SetOmni set a specific omni in the store
func (k Keeper) SetOmni(ctx sdk.Context, omni types.Omni) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OmniKey))
	b := k.cdc.MustMarshal(&omni)
	store.Set(GetOmniIDBytes(omni.Id), b)
}

// GetOmni returns a omni from its id
func (k Keeper) GetOmni(ctx sdk.Context, id uint64) (val types.Omni, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OmniKey))
	b := store.Get(GetOmniIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveOmni removes a omni from the store
func (k Keeper) RemoveOmni(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OmniKey))
	store.Delete(GetOmniIDBytes(id))
}

// GetAllOmni returns all omni
func (k Keeper) GetAllOmni(ctx sdk.Context) (list []types.Omni) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OmniKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Omni
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetOmniIDBytes returns the byte representation of the ID
func GetOmniIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetOmniIDFromBytes returns ID in uint64 format from a byte array
func GetOmniIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
