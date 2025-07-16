package usercmd

import (
	rootCmd "github.com/zakkbob/mxguard/cmd"
	"github.com/spf13/cobra"
	"os"
	"context"
	"github.com/zakkbob/mxguard/cmd/helpers"
	"github.com/zakkbob/mxguard/db"
	"github.com/zakkbob/mxguard/internal/service"
)

var deleteCmd = &cobra.Command {
	Use: "delete",
	Short: "Delete a user",
	Long: "Delete a user",
	Run: func (cmd *cobra.Command, args []string) {
		conn, err := db.InitConn(&rootCmd.Config)
		if err != nil {
			rootCmd.Logger.Fatal().Err(err).Msg("Failed to connect to database")
		}
		repo := db.NewPostgresUserRepository(conn)
		users := service.NewUserService(repo)

		username, err := helpers.GetStringFlagOrPrompt(cmd, os.Stdin, "username", "User to delete: ")
		if err != nil {
			rootCmd.Logger.Fatal().Err(err).Msg("Failed to get username flag")
		}

		//TODO - show user first, to confirm deletion

		err = users.DeleteUserByUsername(context.TODO(), username)
		if err != nil {
			rootCmd.Logger.Fatal().Err(err).Msg("Failed to delete user")
		}
		rootCmd.Logger.Info().Msg("User was deleted")
	},
}

func init() {
	userCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringP("username", "u", "", "Username to delete")
}
