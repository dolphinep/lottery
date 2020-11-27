package lottery

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"gitlab.com/oweliie001/lottery/x/lottery/types"
	"gitlab.com/oweliie001/lottery/x/lottery/keeper"
)

func handleMsgSetTicket(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetTicket) (*sdk.Result, error) {
	var ticket = types.Ticket{
		Creator: msg.Creator,
		ID:      msg.ID,
    	LotteryID: msg.LotteryID,
    	Number: msg.Number,
	}
	if !msg.Creator.Equals(k.GetTicketOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetTicket(ctx, ticket)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
