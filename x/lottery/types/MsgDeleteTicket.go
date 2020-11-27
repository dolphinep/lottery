package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteTicket{}

type MsgDeleteTicket struct {
  ID      string         `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgDeleteTicket(id string, creator sdk.AccAddress) MsgDeleteTicket {
  return MsgDeleteTicket{
    ID: id,
		Creator: creator,
	}
}

func (msg MsgDeleteTicket) Route() string {
  return RouterKey
}

func (msg MsgDeleteTicket) Type() string {
  return "DeleteTicket"
}

func (msg MsgDeleteTicket) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgDeleteTicket) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgDeleteTicket) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}