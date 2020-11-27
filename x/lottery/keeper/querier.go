package keeper

import (
  // this line is used by starport scaffolding # 1
	"gitlab.com/oweliie001/lottery/x/lottery/types"
		
	
		
	
		
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewQuerier creates a new querier for lottery clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
    // this line is used by starport scaffolding # 2
		case types.QueryListPrizeAnnounce:
			return listPrizeAnnounce(ctx, k)
		case types.QueryGetPrizeAnnounce:
			return getPrizeAnnounce(ctx, path[1:], k)
		case types.QueryListTicket:
			return listTicket(ctx, k)
		case types.QueryGetTicket:
			return getTicket(ctx, path[1:], k)
		case types.QueryListTicketById:
			return listTicketByLotteryID(ctx, path[1:], k)
		case types.QueryListLottery:
			return listLottery(ctx, k)
		case types.QueryGetLottery:
			return getLottery(ctx, path[1:], k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown lottery query endpoint")
		}
	}
}
