package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List workspaces",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_listCmd).Standalone()

	help_workspaceCmd.AddCommand(help_workspace_listCmd)
}
