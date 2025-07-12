/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package startcmd

import (
	rootCmd "github.com/zakkbob/mxguard/cmd"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start your server",
	Long:  `start your server`,
	Run: func(cmd *cobra.Command, args []string) {
		rootCmd.Logger.Info().Msg("Preparing to start server")
		rootCmd.Logger.Info().Msg("Listening at ...")
	},
}

func init() {
	rootCmd.RootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
