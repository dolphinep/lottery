package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
    "gitlab.com/oweliie001/lottery/x/lottery/types"
)

func GetCmdListPrizeAnnounce(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "list-prizeAnnounce",
		Short: "list all prizeAnnounce",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/"+types.QueryListPrizeAnnounce, queryRoute), nil)
			if err != nil {
				fmt.Printf("could not list PrizeAnnounce\n%s\n", err.Error())
				return nil
			}
			var out []types.PrizeAnnounce
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdGetPrizeAnnounce(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-prizeAnnounce [key]",
		Short: "Query a prizeAnnounce by key",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			key := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", queryRoute, types.QueryGetPrizeAnnounce, key), nil)
			if err != nil {
				fmt.Printf("could not resolve prizeAnnounce %s \n%s\n", key, err.Error())

				return nil
			}

			var out types.PrizeAnnounce
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
