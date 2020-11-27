package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"gitlab.com/oweliie001/lottery/x/lottery/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	lotteryTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	lotteryTxCmd.AddCommand(flags.PostCommands(
    // this line is used by starport scaffolding # 1
		GetCmdCreatePrizeAnnounce(cdc),
		GetCmdSetPrizeAnnounce(cdc),
		GetCmdDeletePrizeAnnounce(cdc),
		GetCmdCreateTicket(cdc),
		GetCmdSetTicket(cdc),
		GetCmdDeleteTicket(cdc),
		GetCmdCreateLottery(cdc),
		GetCmdSetLottery(cdc),
		GetCmdDeleteLottery(cdc),
	)...)

	return lotteryTxCmd
}
