/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package usercmd

import (
	"github.com/spf13/cobra"
	rootCmd "github.com/zakkbob/mxguard/cmd"
)

// usersCmd represents the users command
var usersCmd = &cobra.Command{
	Use:   "user",
	Short: "Manage users within the application",
	Long:  `M a n a g e   u s e r s   w i t h i n   t h e   a p p l i c a t i o n`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("users called")
	// },
}

func init() {
	rootCmd.RootCmd.AddCommand(usersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// usersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// usersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
