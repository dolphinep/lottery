package lottery

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/crypto"
	"strconv"
	"time"

	"gitlab.com/oweliie001/lottery/x/lottery/keeper"
	"gitlab.com/oweliie001/lottery/x/lottery/types"
)

func handleMsgCreatePrizeAnnounce(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreatePrizeAnnounce) (*sdk.Result, error) {
	var prizeAnnounce = types.PrizeAnnounce{
		Creator:   msg.Creator,
		ID:        msg.ID,
		LotteryID: msg.LotteryID,
	}
	moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))

	lottery, err := k.GetLottery(ctx, msg.LotteryID)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Lottery not exists")
	}

	if lottery.Detail == "announced" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Lottery already announce")
	}

	today := time.Now()
	formatInputTime, err := time.Parse(time.RFC3339, lottery.Drawdate)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Cannot parsed time")
	}
	//check is lottery timeout
	isLotteryAvailable := today.Before(formatInputTime)
	if isLotteryAvailable {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Lottery still available")
	}
	// get all tickets by LotteryID
	ticketList, err := k.ListTicketByLotteryID(ctx, msg.LotteryID)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Some error on finding tickets")
	}
	if len(ticketList) == 0 {
		// TODO: คืนเงินคนตั้งร้าน
		sdkError := k.CoinKeeper.SendCoins(ctx, moduleAcct, lottery.Creator, lottery.Reward)
		if sdkError != nil {
			return nil, sdkError
		}
		prizeAnnounce.FinalReward = "NO ONE BUT THIS LOTTERY!!!"
		// prizeAnnounce.Number = "-"
		//return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "No Tickets Found on "+msg.LotteryID)
	} else {
		// find the final number
		var sum = 0
		for _, ticket := range ticketList {
			for _, c := range ticket.ID {
				sum += (int(c))
			}
		}
		var result = strconv.Itoa(sum % 10)
		prizeAnnounce.Number = result

		//find the lucky ticket
		var winningList []types.Ticket
		for _, tick := range ticketList {
			if tick.Number == result {
				winningList = append(winningList, tick)
			}
		}

		//set prize list of ticket ID
		var winningTickets []string
		for _, x := range winningList {
			winningTickets = append(winningTickets, x.ID)
		}
		prizeAnnounce.TicketID = winningTickets

		if len(winningList) == 0 {
			// TODO: คืนเงินคนตั้งร้าน
			sdkError := k.CoinKeeper.SendCoins(ctx, moduleAcct, lottery.Creator, lottery.Reward)
			if sdkError != nil {
				return nil, sdkError
			}
			prizeAnnounce.FinalReward = "NO WINNER IN THIS ROUND!!!"
		} else {
			// wining >= 1
			var ssf = lottery.Reward.AmountOf("token").Int64() / int64(len(winningList))
			strCoin := strconv.FormatInt(ssf, 10)
			parsedReward, err2 := sdk.ParseCoins(strCoin + "token")
			if err2 != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Parsed Coin Error")
			}
			prizeAnnounce.FinalReward = strCoin + "token"
			//send coin to wining
			for _, winner := range winningList {
				moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
				sdkError := k.CoinKeeper.SendCoins(ctx, moduleAcct, winner.Creator, parsedReward)
				if sdkError != nil {
					return nil, sdkError
				}
			}
		}
		var payback_money = lottery.Price.AmountOf("token").Int64() * int64(len(ticketList))
		strPayback := strconv.FormatInt(payback_money, 10)
		paybackCoin, err2 := sdk.ParseCoins(strPayback + "token")
		if err2 != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Parsed Coin Error")
		}
		// payback money to lottery creator
		sdkError := k.CoinKeeper.SendCoins(ctx, moduleAcct, lottery.Creator, paybackCoin)
		if sdkError != nil {
			return nil, sdkError
		}
	}

	k.CreatePrizeAnnounce(ctx, prizeAnnounce)
	//lottery already announce
	lottery.Detail = "announced"
	k.SetLottery(ctx, lottery)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeCreatePrizeAnnounce),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator.String()),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
