package cmd

import (
	"github.com/spf13/cobra"
	"github.com/srz-zumix/gh-deps-kit/cmd/unity"
)

func init() {
	rootCmd.AddCommand(NewUnityCmd())
}

// NewUnityCmd returns the unity parent command.
func NewUnityCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unity",
		Short: "Manage Unity package dependencies",
	}
	cmd.AddCommand(unity.NewListCmd())
	return cmd
}
