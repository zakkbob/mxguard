/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package userscmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	rootCmd "github.com/zakkbob/mxguard/cmd"
	"github.com/zakkbob/mxguard/internal/database"
	"github.com/zakkbob/mxguard/internal/user"
	"github.com/zakkbob/mxguard/cmd/helpers"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new user",
	Long:  `Create a new user`,
	Run: func(cmd *cobra.Command, args []string) {
		conn := database.Init(&rootCmd.Config)

		var err error
		username, err := helpers.GetStringFlagOrPrompt(cmd, "username", "Enter username: ")
		if err != nil {
			log.WithError(err).Fatal("Failed to get username")
		}

		err = user.CreateUser(conn, username, true)
		if err != nil {
			log.Info(":(")
		}
	},
}

func init() {
	usersCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	createCmd.Flags().String("username", "u", "")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
