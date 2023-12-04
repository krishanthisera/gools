/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package aws

import (
	"fmt"

	"github.com/spf13/cobra"
)

// iamCmd represents the iam command
var iamCmd = &cobra.Command{
	Use:   "iam",
	Short: "AWS IAM Commands",
	Long: `AWS IAM Commands. Use the following commands:
		accesskeyCleanup`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Error: Please specify the IAM Subcommand")
	},
}

func init() {
	awsCmd.AddCommand(iamCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// iamCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// iamCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
