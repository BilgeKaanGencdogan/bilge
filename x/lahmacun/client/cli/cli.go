package cli

import (
	"bilge/x/lahmacun/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

// CmdCreateLahmacun creates a new lahmacun
func CmdCreateLahmacun() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-lahmacun [id] [heat] [dough] [meat] [tomato] [onion]",
		Short: "Create a new lahmacun",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id := args[0]
			heat := args[1]
			dough := args[2]
			meat := args[3]
			tomato := args[4]
			onion := args[5]
			creator := clientCtx.GetFromAddress().String()
			msg := types.NewMsgCreateLahmacun(
				creator,
				id,
				heat,
				dough,
				meat,
				tomato,
				onion,
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

// CmdUpdateLahmacun updates an existing lahmacun
func CmdUpdateLahmacun() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-lahmacun [id] [heat] [dough] [meat] [tomato] [onion]",
		Short: "Update an existing lahmacun",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id := args[0]
			heat := args[1]
			dough := args[2]
			meat := args[3]
			tomato := args[4]
			onion := args[5]
			creator := clientCtx.GetFromAddress().String()
			msg := types.NewMsgCreateLahmacun(creator, id, heat, dough, meat, tomato, onion)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
