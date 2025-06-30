/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/srz-zumix/gh-deps-kit/version"
)

var rootCmd = &cobra.Command{
	Use:     "gh-deps-kit",
	Short:   "A tool to manage GitHub Dependency graph",
	Long:    `gh-deps-kit is a tool to manage GitHub Dependency graph.`,
	Version: version.Version,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
