package helpers

import (
	"fmt"

	"bufio"
	"github.com/spf13/cobra"
	"io"
	"strings"
)

func GetStringFlagOrPrompt(cmd *cobra.Command, reader io.Reader, name string, prompt string) (string, error) {
	if !cmd.Flags().Changed(name) {
		fmt.Print(prompt)
		buf := bufio.NewReader(reader)
		value, err := buf.ReadString('\n')
		if err != nil {
			return "", err
		}
		return strings.TrimRight(value, "\r\n"), nil
	}

	return cmd.Flags().GetString(name)
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
