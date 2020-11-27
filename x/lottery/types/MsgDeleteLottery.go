package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteLottery{}

type MsgDeleteLottery struct {
  ID      string         `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgDeleteLottery(id string, creator sdk.AccAddress) MsgDeleteLottery {
  return MsgDeleteLottery{
    ID: id,
		Creator: creator,
	}
}

func (msg MsgDeleteLottery) Route() string {
  return RouterKey
}

func (msg MsgDeleteLottery) Type() string {
  return "DeleteLottery"
}

func (msg MsgDeleteLottery) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgDeleteLottery) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgDeleteLottery) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}