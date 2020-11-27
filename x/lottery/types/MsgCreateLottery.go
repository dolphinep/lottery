package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateLottery{}

type MsgCreateLottery struct {
	ID       string
	Creator  sdk.AccAddress `json:"creator" yaml:"creator"`
	Name     string         `json:"name" yaml:"name"`
	Detail   string         `json:"detail" yaml:"detail"`
	Reward   sdk.Coins      `json:"reward" yaml:"reward"`
	Drawdate string         `json:"drawdate" yaml:"drawdate"`
	Price    sdk.Coins      `json:"price" yaml:"price"`
}

func NewMsgCreateLottery(creator sdk.AccAddress, name string, detail string, reward sdk.Coins, drawdate string, price sdk.Coins) MsgCreateLottery {
	return MsgCreateLottery{
		ID:       uuid.New().String(),
		Creator:  creator,
		Name:     name,
		Detail:   detail,
		Reward:   reward,
		Drawdate: drawdate,
		Price:    price,
	}
}

func (msg MsgCreateLottery) Route() string {
	return RouterKey
}

func (msg MsgCreateLottery) Type() string {
	return "CreateLottery"
}

func (msg MsgCreateLottery) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateLottery) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgCreateLottery) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
