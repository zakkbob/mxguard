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
