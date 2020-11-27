package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateTicket{}

type MsgCreateTicket struct {
	ID        string
	Creator   sdk.AccAddress `json:"creator" yaml:"creator"`
	LotteryID string         `json:"lotteryID" yaml:"lotteryID"`
	Number    string         `json:"number" yaml:"number"`
}

func NewMsgCreateTicket(creator sdk.AccAddress, lotteryID string, number string) MsgCreateTicket {
	return MsgCreateTicket{
		ID:        uuid.New().String(),
		Creator:   creator,
		LotteryID: lotteryID,
		Number:    number,
	}
}

func (msg MsgCreateTicket) Route() string {
	return RouterKey
}

func (msg MsgCreateTicket) Type() string {
	return "CreateTicket"
}

func (msg MsgCreateTicket) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateTicket) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgCreateTicket) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
