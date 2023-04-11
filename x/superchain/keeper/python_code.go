package keeper

import (
	"encoding/binary"

	"github.com/Peixer/superchain/x/superchain/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetPythonCodeCount get the total number of pythonCode
func (k Keeper) GetPythonCodeCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.PythonCodeCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetPythonCodeCount set the total number of pythonCode
func (k Keeper) SetPythonCodeCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.PythonCodeCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendPythonCode appends a pythonCode in the store with a new id and update the count
func (k Keeper) AppendPythonCode(
	ctx sdk.Context,
	pythonCode types.PythonCode,
) uint64 {
	// Create the pythonCode
	count := k.GetPythonCodeCount(ctx)

	// Set the ID of the appended value
	pythonCode.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PythonCodeKey))
	appendedValue := k.cdc.MustMarshal(&pythonCode)
	store.Set(GetPythonCodeIDBytes(pythonCode.Id), appendedValue)

	// Update pythonCode count
	k.SetPythonCodeCount(ctx, count+1)

	return count
}

// SetPythonCode set a specific pythonCode in the store
func (k Keeper) SetPythonCode(ctx sdk.Context, pythonCode types.PythonCode) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PythonCodeKey))
	b := k.cdc.MustMarshal(&pythonCode)
	store.Set(GetPythonCodeIDBytes(pythonCode.Id), b)
}

// GetPythonCode returns a pythonCode from its id
func (k Keeper) GetPythonCode(ctx sdk.Context, id uint64) (val types.PythonCode, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PythonCodeKey))
	b := store.Get(GetPythonCodeIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePythonCode removes a pythonCode from the store
func (k Keeper) RemovePythonCode(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PythonCodeKey))
	store.Delete(GetPythonCodeIDBytes(id))
}

// GetAllPythonCode returns all pythonCode
func (k Keeper) GetAllPythonCode(ctx sdk.Context) (list []types.PythonCode) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PythonCodeKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PythonCode
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetPythonCodeIDBytes returns the byte representation of the ID
func GetPythonCodeIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetPythonCodeIDFromBytes returns ID in uint64 format from a byte array
func GetPythonCodeIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
