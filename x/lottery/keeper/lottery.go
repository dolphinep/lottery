package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"gitlab.com/oweliie001/lottery/x/lottery/types"
    "github.com/cosmos/cosmos-sdk/codec"
)

// CreateLottery creates a lottery
func (k Keeper) CreateLottery(ctx sdk.Context, lottery types.Lottery) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.LotteryPrefix + lottery.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(lottery)
	store.Set(key, value)
}

// GetLottery returns the lottery information
func (k Keeper) GetLottery(ctx sdk.Context, key string) (types.Lottery, error) {
	store := ctx.KVStore(k.storeKey)
	var lottery types.Lottery
	byteKey := []byte(types.LotteryPrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &lottery)
	if err != nil {
		return lottery, err
	}
	return lottery, nil
}

// SetLottery sets a lottery
func (k Keeper) SetLottery(ctx sdk.Context, lottery types.Lottery) {
	lotteryKey := lottery.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(lottery)
	key := []byte(types.LotteryPrefix + lotteryKey)
	store.Set(key, bz)
}

// DeleteLottery deletes a lottery
func (k Keeper) DeleteLottery(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.LotteryPrefix + key))
}

//
// Functions used by querier
//

func listLottery(ctx sdk.Context, k Keeper) ([]byte, error) {
	var lotteryList []types.Lottery
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.LotteryPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var lottery types.Lottery
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &lottery)
		lotteryList = append(lotteryList, lottery)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, lotteryList)
	return res, nil
}

func getLottery(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	lottery, err := k.GetLottery(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, lottery)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetLotteryOwner(ctx sdk.Context, key string) sdk.AccAddress {
	lottery, err := k.GetLottery(ctx, key)
	if err != nil {
		return nil
	}
	return lottery.Creator
}


// Check if the key exists in the store
func (k Keeper) LotteryExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.LotteryPrefix + key))
}
