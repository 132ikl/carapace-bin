package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var config_help_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Update config file to set the given option to a given value",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_help_setCmd).Standalone()

	config_helpCmd.AddCommand(config_help_setCmd)
}
