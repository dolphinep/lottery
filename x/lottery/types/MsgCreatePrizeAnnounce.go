package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreatePrizeAnnounce{}

type MsgCreatePrizeAnnounce struct {
	ID          string
	Creator     sdk.AccAddress `json:"creator" yaml:"creator"`
	LotteryID   string         `json:"lotteryID" yaml:"lotteryID"`
}

func NewMsgCreatePrizeAnnounce(creator sdk.AccAddress, lotteryID string) MsgCreatePrizeAnnounce {
	return MsgCreatePrizeAnnounce{
		ID:          uuid.New().String(),
		Creator:     creator,
		LotteryID:   lotteryID,
	}
}

func (msg MsgCreatePrizeAnnounce) Route() string {
	return RouterKey
}

func (msg MsgCreatePrizeAnnounce) Type() string {
	return "CreatePrizeAnnounce"
}

func (msg MsgCreatePrizeAnnounce) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreatePrizeAnnounce) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgCreatePrizeAnnounce) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
