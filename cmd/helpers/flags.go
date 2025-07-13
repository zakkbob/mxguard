package helpers

import (
	"fmt"

	"bufio"
	"io"
	"strings"

	"github.com/spf13/cobra"
	"strconv"
)

// Get string cobra flag if set, prompt user if missing
func GetStringFlagOrPrompt(cmd *cobra.Command, reader io.Reader, name string, prompt string) (string, error) {
	if !cmd.Flags().Changed(name) {
		fmt.Print(prompt)
		buf := bufio.NewReader(reader)
		value, err := buf.ReadString('\n')
		if err != nil {
			return "", fmt.Errorf("reading user input: %w", err)
		}
		return strings.TrimRight(value, "\r\n"), nil
	}

	return cmd.Flags().GetString(name)
}

// Get boolean cobra flag if set, prompt user if missing
// Prompt: (Default is capitalised)
// {name} (Y/n):
func GetBoolFlagOrPrompt(cmd *cobra.Command, reader io.Reader, name string) (bool, error) {
	if cmd.Flags().Changed(name) {
		return cmd.Flags().GetBool(name)
	}

	defValue, err := strconv.ParseBool(cmd.Flags().Lookup(name).DefValue)
	if err != nil {
		return false, fmt.Errorf("parsing bool argument's default value: %w", err)
	}
	options := "y/N"
	if defValue {
		options = "Y/n"
	}

	for {
		fmt.Printf("%s (%s): ", name, options)
		buf := bufio.NewReader(reader)
		value, err := buf.ReadString('\n')
		if err != nil {
			return false, fmt.Errorf("reading user input: %w", err)
		}

		value = strings.TrimRight(value, "\r\n")

		if value == "" {
			return defValue, nil
		} else if value == "n" || value == "N" {
			return false, nil
		} else if value == "y" || value == "Y" {
			return true, nil
		}
	}
}

// func MustGetStringFlagOrPrompt(logger zerolog.Logger, cmd *cobra.Command, reader io.Reader, name string, prompt string) string {
// 	if !cmd.Flags().Changed(name) {
// 		fmt.Print(prompt)
// 		buf := bufio.NewReader(reader)
// 		value, err := buf.ReadString('\n')
// 		value = strings.TrimRight(value, "\r\n")
// 		if err != nil {
// 			logger.Fatal().Err(err).Msgf("Failed to get %s", name)
// 		}
// 		return value
// 	}
//
// 	value, err := cmd.Flags().GetString(name)
// 	if err != nil {
// 		logger.Fatal().Err(err).Msgf("Failed to get %s", name)
// 	}
// 	return value
// }
