package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeletePrizeAnnounce{}

type MsgDeletePrizeAnnounce struct {
  ID      string         `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgDeletePrizeAnnounce(id string, creator sdk.AccAddress) MsgDeletePrizeAnnounce {
  return MsgDeletePrizeAnnounce{
    ID: id,
		Creator: creator,
	}
}

func (msg MsgDeletePrizeAnnounce) Route() string {
  return RouterKey
}

func (msg MsgDeletePrizeAnnounce) Type() string {
  return "DeletePrizeAnnounce"
}

func (msg MsgDeletePrizeAnnounce) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgDeletePrizeAnnounce) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgDeletePrizeAnnounce) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}