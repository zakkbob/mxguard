package helpers_test

import (
	"bytes"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/zakkbob/mxguard/cmd/helpers"
	"testing"
)

func TestGetStringFlagOrPrompt_FlagPassed(t *testing.T) {
	var cmd = &cobra.Command{
		Use: "mock",
	}

	cmd.Flags().String("mockflag", "", "mock flag usage")

	cmd.SetArgs([]string{"--mockflag=success"})

	cmd.Execute()

	input := bytes.NewBufferString("fail\n")

	name := "mockflag"
	prompt := ""

	expected := "success"
	got, err := helpers.GetStringFlagOrPrompt(cmd, input, name, prompt)
	assert.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestGetStringFlagOrPrompt_FlagNotPassed(t *testing.T) {
	var cmd = &cobra.Command{
		Use: "mock",
	}

	input := bytes.NewBufferString("success\n")

	name := "mockflag"
	prompt := ""

	expected := "success"
	got, err := helpers.GetStringFlagOrPrompt(cmd, input, name, prompt)
	assert.NoError(t, err)
	assert.Equal(t, expected, got)
}
