package lottery

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"gitlab.com/oweliie001/lottery/x/lottery/types"
	"gitlab.com/oweliie001/lottery/x/lottery/keeper"
)

func handleMsgSetPrizeAnnounce(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetPrizeAnnounce) (*sdk.Result, error) {
	var prizeAnnounce = types.PrizeAnnounce{
		Creator: msg.Creator,
		ID:      msg.ID,
    	LotteryID: msg.LotteryID,
	}
	if !msg.Creator.Equals(k.GetPrizeAnnounceOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetPrizeAnnounce(ctx, prizeAnnounce)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
