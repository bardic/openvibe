package main

import (
	"bytes"
	"testing"
)

var mainTests = []struct {
	args []struct {
		param string
		arg   string
	}
	expected string
}{
	{
		args: []struct {
			param string
			arg   string
		}{
			{param: "", arg: ""},
		},
		expected: "openvibe -- v0.0.1",
	},
}

func TestRoot(t *testing.T) {

}

func TestVersion(t *testing.T) {
	actual := new(bytes.Buffer)
	cmdArgs := []string{"openvibe", "version"}

	Cmd.Root().SetOut(actual)
	Cmd.Root().SetErr(actual)

	for _, test := range mainTests {
		for _, args := range test.args {
			cmdArgs = append(cmdArgs, args.param, args.arg)
		}

		Cmd.Root().SetArgs(cmdArgs)
		Cmd.Root().Execute()
	}
}
