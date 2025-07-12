/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package userscmd

import (
	"github.com/spf13/cobra"
	rootCmd "github.com/zakkbob/mxguard/cmd"
	"github.com/zakkbob/mxguard/cmd/helpers"
	"github.com/zakkbob/mxguard/internal/database"
	"github.com/zakkbob/mxguard/db"
	"os"
	"context"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new user",
	Long:  `Create a new user`,
	Run: func(cmd *cobra.Command, args []string) {
		conn := database.Init(rootCmd.Logger, &rootCmd.Config)
		userRepository := db.NewPostgresUserRepository(conn)

		var err error
		username, err := helpers.GetStringFlagOrPrompt(cmd, os.Stdin, "username", "Enter username: ")
		if err != nil {
			rootCmd.Logger.Fatal().Err(err).Msg("Failed to get username")
		}

		user, err := userRepository.CreateUser(context.TODO(), username, true)
		if err != nil {
			rootCmd.Logger.Fatal().Err(err).Str("username", username).Bool("isAdmin", true).Msg("Failed to create user")
		}
		rootCmd.Logger.Info().Any("user", user).Msg("Successfully created user")
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
