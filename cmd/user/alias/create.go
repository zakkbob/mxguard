/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package aliascmd

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	rootcmd "github.com/zakkbob/mxguard/cmd"
	"github.com/zakkbob/mxguard/cmd/helpers"
	"github.com/zakkbob/mxguard/db"
)

// createCmd represents the created command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a user alias",
	Long: `Create a user alias
	Example: mxguard user alias created --user="John Doe" --alias="abc123" --description="Testing"`,
	Run: func(cmd *cobra.Command, args []string) {
		username, err := helpers.GetStringFlagOrPrompt(cmd, os.Stdin, "user", "Enter username: ")
		if err != nil {
			rootcmd.Logger.Fatal().Err(err).Msg("Failed to get username argument")
		}

		alias, err := helpers.GetStringFlagOrPrompt(cmd, os.Stdin, "alias", "Enter alias name: ")
		if err != nil {
			rootcmd.Logger.Fatal().Err(err).Msg("Failed to get alias argument")
		}

		description, err := helpers.GetStringFlagOrPrompt(cmd, os.Stdin, "description", "Enter alias description: ")
		if err != nil {
			rootcmd.Logger.Fatal().Err(err).Msg("Failed to get description argument")
		}

		conn, err := db.InitConn(&rootcmd.Config)
		if err != nil {
			rootcmd.Logger.Fatal().Err(err).Msg("Failed to connect to database")
		}
		defer conn.Close()

		userRepo := db.NewPostgresUserRepository(conn)

		user, err := userRepo.GetUserByUsername(context.TODO(), username)
		if err != nil {
			rootcmd.Logger.Fatal().Err(err).Msg("Failed to get user")
		}

		_, err = userRepo.CreateAlias(context.Background(), user, alias, description)
		if err != nil {
			rootcmd.Logger.Fatal().Err(err).Msg("Failed to create alias")
		}
		rootcmd.Logger.Info().Msg("Successfully created alias")
	},
}

func init() {
	aliasCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// usersCmd.PersistentFlags().String("foo", "", "A help for foo")
	createCmd.Flags().StringP("user", "", "", "Username of user")
	createCmd.Flags().StringP("alias", "", "", "Set alias name")
	createCmd.Flags().StringP("description", "", "", "Set alias description")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// usersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
