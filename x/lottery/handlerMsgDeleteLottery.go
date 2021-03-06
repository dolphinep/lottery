package lottery

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"gitlab.com/oweliie001/lottery/x/lottery/types"
	"gitlab.com/oweliie001/lottery/x/lottery/keeper"
)

// Handle a message to delete name
func handleMsgDeleteLottery(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeleteLottery) (*sdk.Result, error) {
	if !k.LotteryExists(ctx, msg.ID) {
		// replace with ErrKeyNotFound for 0.39+
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.ID)
	}
	if !msg.Creator.Equals(k.GetLotteryOwner(ctx, msg.ID)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	k.DeleteLottery(ctx, msg.ID)
	return &sdk.Result{}, nil
}
