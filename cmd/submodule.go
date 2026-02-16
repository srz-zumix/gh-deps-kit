package cmd

import (
	"github.com/spf13/cobra"
	"github.com/srz-zumix/gh-deps-kit/cmd/submodule"
)

func init() {
	rootCmd.AddCommand(NewSubmoduleCmd())
}

// NewSubmoduleCmd returns the submodule parent command
func NewSubmoduleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submodule",
		Short: "Manage repository submodules",
	}
	cmd.AddCommand(submodule.NewListCmd())
	return cmd
}
