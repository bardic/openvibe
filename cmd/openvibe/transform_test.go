package main

import (
	"bytes"
	"testing"
)

var transformTests = []struct {
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
			{param: "--target-assets", arg: "blocks"},
			{param: "--in", arg: "./test_data/sample"},
			//{param: "--out", arg: "/tmp"},
			{param: "--capability", arg: "pbr"},
			{param: "--modifiers", arg: "test.png test.jpg"},
		},
		expected: "openvibe -- v0.0.1",
	},
}

func TestTransform(t *testing.T) {
	actual := new(bytes.Buffer)
	cmdArgs := []string{"openvibe", "version"}

	Cmd.Root().SetOut(actual)
	Cmd.Root().SetErr(actual)

	for _, test := range transformTests {
		for _, args := range test.args {
			cmdArgs = append(cmdArgs, args.param, args.arg)
		}

		Cmd.Root().SetArgs(cmdArgs)
		Cmd.Root().Execute()
	}
}
