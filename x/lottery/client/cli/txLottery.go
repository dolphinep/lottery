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

func GetCmdCreateLottery(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-lottery [name] [detail] [reward] [drawdate] [price]",
		Short: "Creates a new lottery",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsName := string(args[0] )
			argsDetail := string(args[1] )
			argsReward := string(args[2] )
			argsDrawdate := string(args[3] )
			argsPrice := string(args[4] )
			
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			reward, err := sdk.ParseCoins(argsReward)
			if err != nil {
				return err
			}
			price, err := sdk.ParseCoins(argsPrice)
			if err != nil {
				return err
			}
			msg := types.NewMsgCreateLottery(cliCtx.GetFromAddress(), string(argsName), string(argsDetail), reward, string(argsDrawdate), price)
			err2 := msg.ValidateBasic()
			if err2 != nil {
				return err2
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}


func GetCmdSetLottery(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set-lottery [id]  [name] [detail] [reward] [drawdate] [price]",
		Short: "Set a new lottery",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			argsName := string(args[1])
			argsDetail := string(args[2])
			argsReward := string(args[3])
			argsDrawdate := string(args[4])
			argsPrice := string(args[5])
			
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			reward, err := sdk.ParseCoins(argsReward)
			if err != nil {
				return err
			}
			price, err := sdk.ParseCoins(argsPrice)
			if err != nil {
				return err
			}
			msg := types.NewMsgSetLottery(cliCtx.GetFromAddress(), id, string(argsName), string(argsDetail), reward, string(argsDrawdate), price)
			err2 := msg.ValidateBasic()
			if err2 != nil {
				return err2
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdDeleteLottery(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-lottery [id]",
		Short: "Delete a new lottery by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgDeleteLottery(args[0], cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
