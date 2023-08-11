package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"omni/x/omni/types"
)

var _ = strconv.Itoa(0)

func CmdGetStorage() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-storage [address] [position] [block-tag]",
		Short: "Broadcast message get-storage",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAddress := args[0]
			argPosition := args[1]
			argBlockTag := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgGetStorage(
				clientCtx.GetFromAddress().String(),
				argAddress,
				argPosition,
				argBlockTag,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
