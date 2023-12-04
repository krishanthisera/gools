/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package aws

import (
	"fmt"

	"github.com/spf13/cobra"
)

// usersCmd represents the users command
var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "This would list all the users in the AWS account",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("users called")
	},
}

func init() {
	iamCmd.AddCommand(usersCmd)

}
