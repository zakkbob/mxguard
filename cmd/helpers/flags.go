package helpers

import (
	"fmt"

	"github.com/spf13/cobra"
)

func GetStringFlagOrPrompt(cmd *cobra.Command, name string, prompt string) (string, error) {
	var value string

	if !cmd.Flags().Changed(name) {
		fmt.Print(prompt)
		fmt.Scanln(&value)
		return value, nil
	}

	value, err := cmd.Flags().GetString(name)
	return "", err
}
