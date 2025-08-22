package main

import (
	"os"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

type GenCmd struct {
	// OVCommand
}

type WorkspaceStruct struct {
	Version int      `json:"v"`
	Folders []string `json:"folders"`
}

func newGenCmd(fs afero.Fs) *cobra.Command {
	return &cobra.Command{
		Use:   "gen",
		Short: "generates new workspace filestructure",
		RunE: func(cmd *cobra.Command, args []string) error {
			return os.RemoveAll(args[0])
		},
	}
}

// var (
// 	FS  afero.Fs
// 	Cmd = &cobra.Command{
// 		Use:   "gen",
// 		Short: "generates new workspace filestructure",
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			AppFs := afero.NewMemMapFs()

// 			conf := "./assets/configs/workspace/struct.json"
// 			jsonFile, err := AppFs.Open(conf)
// 			//jsonFile, err := os.Open(conf)
// 			if err != nil {
// 				return err
// 			}

// 			defer jsonFile.Close()

// 			byteValue, err := io.ReadAll(jsonFile)
// 			if err != nil {
// 				return err
// 			}

// 			var jsonConfig WorkspaceStruct
// 			err = json.Unmarshal(byteValue, &jsonConfig)
// 			if err != nil {
// 				return err
// 			}

// 			if jsonConfig.Version != 1 {
// 				return errors.New("meow")
// 			}

// 			for _, f := range jsonConfig.Folders {
// 				os.MkdirAll(f, os.ModeDir)
// 			}

// 			return nil
// 		},
// 	}
// )

// func init() {

// }
