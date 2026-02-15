package cmd

import (
	"github.com/spf13/cobra"
	"github.com/srz-zumix/gh-deps-kit/cmd/actions"
)

func init() {
	rootCmd.AddCommand(NewActionsCmd())
}

// NewActionsCmd returns the actions parent command
func NewActionsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "actions",
		Short: "Manage GitHub Actions-related dependencies",
	}
	cmd.AddCommand(actions.NewListCmd())
	cmd.AddCommand(actions.NewGraphCmd())
	return cmd
}
