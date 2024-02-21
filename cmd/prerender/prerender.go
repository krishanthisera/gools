/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package prerender

import (
	"fmt"

	"github.com/spf13/cobra"

	goolsCmd "krishanthisera/gools/cmd"
)

// prerenderCmd represents the prerender command
var prerenderCmd = &cobra.Command{
	Use:   "prerender",
	Short: "This can be used to debug prerender issues",
	Long:  `This can be used to debug prerender issues. For example: check if the prerender server is returning the correct status code for a given URL.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("prerender called")
	},
}

func init() {
	goolsCmd.RootCmd.AddCommand(prerenderCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// prerenderCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// prerenderCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
