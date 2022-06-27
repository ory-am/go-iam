package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ory/hydra/cmd/cliclient"
	"github.com/ory/x/cmdx"
)

func NewUpdateCmd(root *cobra.Command) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "update",
		Short: "Update resources",
	}
	cmd.AddCommand(NewUpdateClientCmd(root))
	cliclient.RegisterClientFlags(cmd.PersistentFlags())
	cmdx.RegisterFormatFlags(cmd.PersistentFlags())
	return cmd
}
