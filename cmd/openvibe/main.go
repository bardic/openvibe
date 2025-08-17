package main

import (
	"github.com/spf13/afero"
)

func main() {
	appFs := afero.NewOsFs()

	err := NewRootCmd(appFs).Execute()
	if err != nil {
		panic("Error: " + err.Error())
	}
}
