package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Lottery struct {
	Creator  sdk.AccAddress `json:"creator" yaml:"creator"`
	ID       string         `json:"id" yaml:"id"`
	Name     string         `json:"name" yaml:"name"`
	Detail   string         `json:"detail" yaml:"detail"`
	Reward   sdk.Coins      `json:"reward" yaml:"reward"`
	Drawdate string         `json:"drawdate" yaml:"drawdate"`
	Price    sdk.Coins      `json:"price" yaml:"price"`
}
