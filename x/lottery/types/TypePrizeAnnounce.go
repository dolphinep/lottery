package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type PrizeAnnounce struct {
	Creator     sdk.AccAddress `json:"creator" yaml:"creator"`
	ID          string         `json:"id" yaml:"id"`
	LotteryID   string         `json:"lotteryID" yaml:"lotteryID"`
	TicketID    []string         `json:"ticketID" yaml:"ticketID"`
	Number      string         `json:"number" yaml:"number"`
	FinalReward string         `json:"finalReward" yaml:"finalReward"`
}
