package lottery

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/oweliie001/lottery/x/lottery/keeper"
	"gitlab.com/oweliie001/lottery/x/lottery/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
    // this line is used by starport scaffolding # 1
		case types.MsgCreatePrizeAnnounce:
			return handleMsgCreatePrizeAnnounce(ctx, k, msg)
		case types.MsgSetPrizeAnnounce:
			return handleMsgSetPrizeAnnounce(ctx, k, msg)
		case types.MsgDeletePrizeAnnounce:
			return handleMsgDeletePrizeAnnounce(ctx, k, msg)
		case types.MsgCreateTicket:
			return handleMsgCreateTicket(ctx, k, msg)
		case types.MsgSetTicket:
			return handleMsgSetTicket(ctx, k, msg)
		case types.MsgDeleteTicket:
			return handleMsgDeleteTicket(ctx, k, msg)
		case types.MsgCreateLottery:
			return handleMsgCreateLottery(ctx, k, msg)
		case types.MsgSetLottery:
			return handleMsgSetLottery(ctx, k, msg)
		case types.MsgDeleteLottery:
			return handleMsgDeleteLottery(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
