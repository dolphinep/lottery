package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetLottery{}

type MsgSetLottery struct {
  ID      string      `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  Name string `json:"name" yaml:"name"`
  Detail string `json:"detail" yaml:"detail"`
  Reward sdk.Coins `json:"reward" yaml:"reward"`
  Drawdate string `json:"drawdate" yaml:"drawdate"`
  Price sdk.Coins `json:"price" yaml:"price"`
}

func NewMsgSetLottery(creator sdk.AccAddress, id string, name string, detail string, reward sdk.Coins, drawdate string, price sdk.Coins) MsgSetLottery {
  return MsgSetLottery{
    ID: id,
		Creator: creator,
    Name: name,
    Detail: detail,
    Reward: reward,
    Drawdate: drawdate,
    Price: price,
	}
}

func (msg MsgSetLottery) Route() string {
  return RouterKey
}

func (msg MsgSetLottery) Type() string {
  return "SetLottery"
}

func (msg MsgSetLottery) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgSetLottery) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgSetLottery) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}