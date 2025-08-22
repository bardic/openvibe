package main

import (
	"fmt"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

func main() {
	appFs := afero.NewOsFs()

	err := newRootCmd(appFs).Execute()
	if err != nil {
		panic("Error: " + err.Error())
	}
}

func newRootCmd(fs afero.Fs) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "openvibe",
		Short: "Create Vibrant Textures Packs for Minecraft Bedrock",
	}

	cmd.AddCommand(newDownloadCmd(fs))
	cmd.AddCommand(newTransformCmd(fs))
	cmd.AddCommand(newBuildCmd(fs))
	cmd.AddCommand(newGenCmd(fs))
	cmd.AddCommand(NewVersionCmd())

	return cmd
}

func NewVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use: "version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("openvibe -- v0.0.1")
		},
	}
}
