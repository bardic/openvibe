package main

import (
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

func NewRootCmd(fs afero.Fs) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "openvibe",
		Short: "Create Vibrant Textures Packs for Minecraft Bedrock",
	}

	cmd.AddCommand(NewDownloadCmd(fs))
	cmd.AddCommand(NewTransformCmd(fs))
	cmd.AddCommand(NewVersionCmd(fs))

	return cmd
}
