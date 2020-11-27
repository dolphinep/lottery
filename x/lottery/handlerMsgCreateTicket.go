package lottery

import (
	"strconv"
	"time"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/crypto"

	"gitlab.com/oweliie001/lottery/x/lottery/types"
	"gitlab.com/oweliie001/lottery/x/lottery/keeper"
)


func handleMsgCreateTicket(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateTicket) (*sdk.Result, error) {
	var ticket = types.Ticket{
		Creator:   msg.Creator,
		ID:        msg.ID,
		LotteryID: msg.LotteryID,
		Number: msg.Number,
	}
	lottery, err := k.GetLottery(ctx, msg.LotteryID)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Lottery not exists")
	}
	//check announced lottery
	if  lottery.Detail == "announced" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Lottery already announce")
	}
	today := time.Now()
	formatInputTime, err := time.Parse(time.RFC3339, lottery.Drawdate)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Cannot parsed time")
	}
	//check is lottery timeout
	isLotteryAvailable := today.After(formatInputTime)
	if isLotteryAvailable {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Cannot buy tickets: Lottery are available to announce")
	}
	//check input number
	numInt, err := strconv.Atoi(ticket.Number)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Please Input Number")
	}
	if numInt >= 10 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Please buy ticket number: 0-9")
	}
	moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	sdkError := k.CoinKeeper.SendCoins(ctx, ticket.Creator, moduleAcct, lottery.Price)
	if sdkError != nil {
		return nil, sdkError
	}
	k.CreateTicket(ctx, ticket)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeCreateTicket),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator.String()),
			sdk.NewAttribute(types.AttributeLottery, msg.LotteryID),
			sdk.NewAttribute(types.AttributeNumber, msg.Number),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
