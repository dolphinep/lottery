package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	// "github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"gitlab.com/oweliie001/lottery/x/lottery/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	// Group lottery queries under a subcommand
	lotteryQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	lotteryQueryCmd.AddCommand(
		flags.GetCommands(
      // this line is used by starport scaffolding # 1
			GetCmdListPrizeAnnounce(queryRoute, cdc),
			GetCmdGetPrizeAnnounce(queryRoute, cdc),
			GetCmdListTicket(queryRoute, cdc),
			GetCmdGetTicket(queryRoute, cdc),
			GetCmdListLottery(queryRoute, cdc),
			GetCmdGetLottery(queryRoute, cdc),
		)...
	)

	return lotteryQueryCmd
}
