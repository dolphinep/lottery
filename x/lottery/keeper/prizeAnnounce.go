package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"gitlab.com/oweliie001/lottery/x/lottery/types"
    "github.com/cosmos/cosmos-sdk/codec"
)

// CreatePrizeAnnounce creates a prizeAnnounce
func (k Keeper) CreatePrizeAnnounce(ctx sdk.Context, prizeAnnounce types.PrizeAnnounce) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.PrizeAnnouncePrefix + prizeAnnounce.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(prizeAnnounce)
	store.Set(key, value)
}

// GetPrizeAnnounce returns the prizeAnnounce information
func (k Keeper) GetPrizeAnnounce(ctx sdk.Context, key string) (types.PrizeAnnounce, error) {
	store := ctx.KVStore(k.storeKey)
	var prizeAnnounce types.PrizeAnnounce
	byteKey := []byte(types.PrizeAnnouncePrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &prizeAnnounce)
	if err != nil {
		return prizeAnnounce, err
	}
	return prizeAnnounce, nil
}

// SetPrizeAnnounce sets a prizeAnnounce
func (k Keeper) SetPrizeAnnounce(ctx sdk.Context, prizeAnnounce types.PrizeAnnounce) {
	prizeAnnounceKey := prizeAnnounce.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(prizeAnnounce)
	key := []byte(types.PrizeAnnouncePrefix + prizeAnnounceKey)
	store.Set(key, bz)
}

// DeletePrizeAnnounce deletes a prizeAnnounce
func (k Keeper) DeletePrizeAnnounce(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.PrizeAnnouncePrefix + key))
}

//
// Functions used by querier
//

func listPrizeAnnounce(ctx sdk.Context, k Keeper) ([]byte, error) {
	var prizeAnnounceList []types.PrizeAnnounce
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.PrizeAnnouncePrefix))
	for ; iterator.Valid(); iterator.Next() {
		var prizeAnnounce types.PrizeAnnounce
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &prizeAnnounce)
		prizeAnnounceList = append(prizeAnnounceList, prizeAnnounce)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, prizeAnnounceList)
	return res, nil
}

func getPrizeAnnounce(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	prizeAnnounce, err := k.GetPrizeAnnounce(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, prizeAnnounce)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetPrizeAnnounceOwner(ctx sdk.Context, key string) sdk.AccAddress {
	prizeAnnounce, err := k.GetPrizeAnnounce(ctx, key)
	if err != nil {
		return nil
	}
	return prizeAnnounce.Creator
}


// Check if the key exists in the store
func (k Keeper) PrizeAnnounceExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.PrizeAnnouncePrefix + key))
}
