package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetPrizeAnnounce{}

type MsgSetPrizeAnnounce struct {
  ID      string      `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  LotteryID string `json:"lotteryID" yaml:"lotteryID"`
  TicketID string `json:"ticketID" yaml:"ticketID"`
  Number string `json:"number" yaml:"number"`
  FinalReward string `json:"finalReward" yaml:"finalReward"`
}

func NewMsgSetPrizeAnnounce(creator sdk.AccAddress, id string, lotteryID string, ticketID string, number string, finalReward string) MsgSetPrizeAnnounce {
  return MsgSetPrizeAnnounce{
    ID: id,
		Creator: creator,
    LotteryID: lotteryID,
    TicketID: ticketID,
    Number: number,
    FinalReward: finalReward,
	}
}

func (msg MsgSetPrizeAnnounce) Route() string {
  return RouterKey
}

func (msg MsgSetPrizeAnnounce) Type() string {
  return "SetPrizeAnnounce"
}

func (msg MsgSetPrizeAnnounce) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgSetPrizeAnnounce) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgSetPrizeAnnounce) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}