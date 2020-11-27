package cli

import (
	"bufio"
    
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"gitlab.com/oweliie001/lottery/x/lottery/types"
)

func GetCmdCreatePrizeAnnounce(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-prizeAnnounce [lotteryID] [ticketID] [number] [finalReward]",
		Short: "Creates a new prizeAnnounce",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsLotteryID := string(args[0] )
			// argsTicketID := string(args[1] )
			// argsNumber := string(args[2] )
			// argsFinalReward := string(args[3] )
			
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreatePrizeAnnounce(cliCtx.GetFromAddress(), string(argsLotteryID))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}


func GetCmdSetPrizeAnnounce(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set-prizeAnnounce [id]  [lotteryID] [ticketID] [number] [finalReward]",
		Short: "Set a new prizeAnnounce",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			argsLotteryID := string(args[1])
			argsTicketID := string(args[2])
			argsNumber := string(args[3])
			argsFinalReward := string(args[4])
			
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgSetPrizeAnnounce(cliCtx.GetFromAddress(), id, string(argsLotteryID), string(argsTicketID), string(argsNumber), string(argsFinalReward))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdDeletePrizeAnnounce(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-prizeAnnounce [id]",
		Short: "Delete a new prizeAnnounce by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgDeletePrizeAnnounce(args[0], cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
