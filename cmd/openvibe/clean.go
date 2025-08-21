package main

import (
	"os"

	"github.com/spf13/cobra"
)

func cleanCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "clean",
		Short: "removes all generated files from workspace",
		RunE: func(cmd *cobra.Command, args []string) error {
			return os.RemoveAll(args[0])
		},
	}
}
