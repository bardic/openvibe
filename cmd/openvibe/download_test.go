package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

var cleanTests = []struct {
	in  string
	out string
}{
	{"%a", "[%a]"},
	{"%-a", "[%-a]"},
}

func TestDownloand(t *testing.T) {
	baseDir := "./base"
	downloadDir := "./downloads"

	fs := afero.NewMemMapFs()
	cmd := NewDownloadCmd(fs)
	actual := new(bytes.Buffer)
	cmd.Root().SetOut(actual)
	cmd.Root().SetErr(actual)
	cmd.Root().SetArgs([]string{"download", "--url", "https://api.github.com/repos/Mojang/bedrock-samples/releases/latest"})
	cmd.Root().Execute()

	expected := ""

	assert.Equal(t, actual.String(), expected, "actual is not expected")

	_, err := os.Stat(filepath.Join(downloadDir, "latest.zip"))
	assert.False(t, os.IsNotExist(err), "File should exist")

	_, err = os.Stat(filepath.Join(baseDir, "Mojang-bedrock-samples-*"))
	assert.False(t, os.IsNotExist(err), "File should exist")

	err = os.RemoveAll(downloadDir)
	assert.ErrorIs(t, err, nil)

	err = os.RemoveAll(baseDir)
	assert.ErrorIs(t, err, nil)
}
