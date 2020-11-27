package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetTicket{}

type MsgSetTicket struct {
  ID      string      `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  LotteryID string `json:"lotteryID" yaml:"lotteryID"`
  Number string `json:"number" yaml:"number"`
}

func NewMsgSetTicket(creator sdk.AccAddress, id string, lotteryID string, number string) MsgSetTicket {
  return MsgSetTicket{
    ID: id,
		Creator: creator,
    LotteryID: lotteryID,
    Number: number,
	}
}

func (msg MsgSetTicket) Route() string {
  return RouterKey
}

func (msg MsgSetTicket) Type() string {
  return "SetTicket"
}

func (msg MsgSetTicket) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgSetTicket) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgSetTicket) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}