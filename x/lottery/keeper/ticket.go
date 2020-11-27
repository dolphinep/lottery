package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"gitlab.com/oweliie001/lottery/x/lottery/types"
    "github.com/cosmos/cosmos-sdk/codec"
)

// CreateTicket creates a ticket
func (k Keeper) CreateTicket(ctx sdk.Context, ticket types.Ticket) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.TicketPrefix + ticket.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(ticket)
	store.Set(key, value)
}

// GetTicket returns the ticket information
func (k Keeper) GetTicket(ctx sdk.Context, key string) (types.Ticket, error) {
	store := ctx.KVStore(k.storeKey)
	var ticket types.Ticket
	byteKey := []byte(types.TicketPrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &ticket)
	if err != nil {
		return ticket, err
	}
	return ticket, nil
}

// SetTicket sets a ticket
func (k Keeper) SetTicket(ctx sdk.Context, ticket types.Ticket) {
	ticketKey := ticket.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(ticket)
	key := []byte(types.TicketPrefix + ticketKey)
	store.Set(key, bz)
}

// DeleteTicket deletes a ticket
func (k Keeper) DeleteTicket(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.TicketPrefix + key))
}

//
// Functions used by querier
//

// get ticket by lotteryID
func (k Keeper) ListTicketByLotteryID(ctx sdk.Context, lotteryID string) ([]types.Ticket, error) {
	var ticketList []types.Ticket
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.TicketPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var ticket types.Ticket
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &ticket)
		if ticket.LotteryID == lotteryID {
			ticketList = append(ticketList, ticket)
		}
	}
	return ticketList, nil
}
// TODO: TOP
func listTicketByLotteryID(ctx sdk.Context, path []string,k Keeper) ([]byte, error) {
	var ticketList []types.Ticket
	key := path[0]
	ticketList, err := k.ListTicketByLotteryID(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err := codec.MarshalJSONIndent(k.cdc, ticketList)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func listTicket(ctx sdk.Context, k Keeper) ([]byte, error) {
	var ticketList []types.Ticket
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.TicketPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var ticket types.Ticket
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &ticket)
		ticketList = append(ticketList, ticket)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, ticketList)
	return res, nil
}

func getTicket(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	ticket, err := k.GetTicket(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, ticket)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetTicketOwner(ctx sdk.Context, key string) sdk.AccAddress {
	ticket, err := k.GetTicket(ctx, key)
	if err != nil {
		return nil
	}
	return ticket.Creator
}


// Check if the key exists in the store
func (k Keeper) TicketExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.TicketPrefix + key))
}
