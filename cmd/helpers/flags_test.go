package helpers_test

import (
	"fmt"
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

	cmd.Flags().String("mockflag", "", "mock flag usage")
	cmd.Execute()

	input := bytes.NewBufferString("success\n")

	name := "mockflag"
	prompt := ""

	expected := "success"
	got, err := helpers.GetStringFlagOrPrompt(cmd, input, name, prompt)

	assert.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestGetBoolFlagOrPrompt(t *testing.T) {
	tests := []struct{
		input string
		flag string
		defValue bool
		expected bool
	}{
		{"n", "", false, false},
		{"N", "", false, false},
		{"y", "", false, true},
		{"Y", "", false, true},
		{"", "", false, false},
		{"", "", true, true},
		{"", "true", false, true},
		{"", "false", false, false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s_%s_%t", tt.input, tt.flag, tt.defValue), func(t *testing.T) {
			var cmd = &cobra.Command{
				Use: "mock",
			}
			name := "mockflag"

			cmd.Flags().Bool(name, tt.defValue, "mock flag usage")
			if tt.flag != "" {
				cmd.SetArgs([]string{fmt.Sprintf("--mockflag=%s", tt.flag)})
			}
			cmd.Execute()

			input := bytes.NewBufferString(tt.input + string('\n'))

			got, err := helpers.GetBoolFlagOrPrompt(cmd, input, name)

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, got)
		})
	}
}
