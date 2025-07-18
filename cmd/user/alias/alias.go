/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package aliascmd

import (
	"github.com/spf13/cobra"
	usercmd "github.com/zakkbob/mxguard/cmd/user"
)

// alisCmd represents the alias command
var aliasCmd = &cobra.Command{
	Use:   "alias",
	Short: "Manage user aliases",
	Long:  `Manage user aliases`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("users called")
	// },
}

func init() {
	usercmd.UserCmd.AddCommand(aliasCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// usersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// usersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
