package lottery

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/crypto"

	"gitlab.com/oweliie001/lottery/x/lottery/keeper"
	"gitlab.com/oweliie001/lottery/x/lottery/types"
)

func handleMsgCreateLottery(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateLottery) (*sdk.Result, error) {
	var lottery = types.Lottery{
		Creator:  msg.Creator,
		ID:       msg.ID,
		Name:     msg.Name,
		Detail:   msg.Detail,
		Reward:   msg.Reward,
		Drawdate: msg.Drawdate,
		Price:    msg.Price,
	}
	_, err := k.GetLottery(ctx, msg.ID)
	if err == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Lottery already exists")
	}
	moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	sdkError := k.CoinKeeper.SendCoins(ctx, lottery.Creator, moduleAcct, lottery.Reward)
	if sdkError != nil {
		return nil, sdkError
	}

	k.CreateLottery(ctx, lottery)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeCreateLottery),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator.String()),
			sdk.NewAttribute(types.AttributeDetail, msg.Detail),
			sdk.NewAttribute(types.AttributeReward, msg.Reward.String()),
			sdk.NewAttribute(types.AttributePrice, msg.Price.String()),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
