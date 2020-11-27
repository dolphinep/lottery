package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Ticket struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
    LotteryID string `json:"lotteryID" yaml:"lotteryID"`
    Number string `json:"number" yaml:"number"`
}