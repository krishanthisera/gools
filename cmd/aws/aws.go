/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package aws

import (
	gools_cmd "krishanthisera/gools/cmd"

	"github.com/spf13/cobra"
)

var awsProfile string
var awsRegion string

// awsCmd represents the aws command
var awsCmd = &cobra.Command{
	Use:   "aws",
	Short: "AWS Cloud Management Commands",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	gools_cmd.RootCmd.AddCommand(awsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	awsCmd.PersistentFlags().String("profile", "", "AWS Profile name as you have define in ~/.aws/credentials or ~/.aws/config")
	awsCmd.PersistentFlags().String("region", "ap-southeast-2", "AWS Region")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// awsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
