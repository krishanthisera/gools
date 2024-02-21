/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package prerender

import (
	goolsCmd "krishanthisera/gools/cmd"

	"github.com/spf13/cobra"
)

var (
	baseUrls   []string
	subPaths   []string
	pdp        []string
	invalidPdp string
	userAgent  string
)

// testCrawlingCmd represents the testCrawling command
var testCrawlingCmd = &cobra.Command{
	Use:   "testCrawling",
	Short: "Crawl through given set of urls and check if the status code is as expected",
	Long:  `Crawl through given set of urls and check if the status code is as expected. For example: check if the prerender server is returning the correct status code for a given URL.`,
	Run: func(cmd *cobra.Command, args []string) {
		baseUrls, _ = cmd.Flags().GetStringArray("baseUrls")
		subPaths, _ = cmd.Flags().GetStringArray("subPaths")
		pdp, _ = cmd.Flags().GetStringArray("pdp")
		invalidPdp, _ = cmd.Flags().GetString("invalidPdp")
		userAgent, _ = cmd.Flags().GetString("userAgent")

	},
}

func init() {

	testCrawlingCmd.PersistentFlags().StringArray("baseUrls", nil, "set of base urls to crawl")
	testCrawlingCmd.PersistentFlags().StringArray("subPaths", nil, "sub-paths to append to the base URL")
	testCrawlingCmd.PersistentFlags().StringArray("pdp", nil, "working set of PDPs")
	testCrawlingCmd.PersistentFlags().String("invalidPdp", "this-is-not-a-pdp", "invalid PDP to test")
	testCrawlingCmd.PersistentFlags().String("userAgent", "Googlebot", "user agent to use for crawling")

	goolsCmd.RootCmd.AddCommand(testCrawlingCmd)

}
