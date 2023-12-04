/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package aws

import (
	"fmt"
	goolsCmd "krishanthisera/gools/cmd"
	goolsIam "krishanthisera/gools/lib/aws/iam"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/spf13/cobra"
)

// keyCleanupCmd represents the keyCleanup command
var keyCleanupCmd = &cobra.Command{
	Use:   "keyCleanup",
	Short: "This would deactivate the IAM access key AWS account wide",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		awsProfile, _ = cmd.Flags().GetString("profile")
		awsRegion, _ = cmd.Flags().GetString("region")

		if !goolsCmd.ConfirmAction() {
			fmt.Println("Exiting...")
			return
		}

		opts := session.Options{}

		// If the profile is provided, use it
		if awsProfile != "" {
			fmt.Println("Using AWS Profile:", awsProfile)
			opts = session.Options{
				Profile:                 awsProfile,
				SharedConfigState:       session.SharedConfigEnable,
				AssumeRoleTokenProvider: stscreds.StdinTokenProvider,
				Config: aws.Config{
					Region: aws.String(awsRegion), // Replace with your desired AWS region
				},
			}
		} else {
			opts = session.Options{
				Config: aws.Config{
					Region: aws.String(awsRegion), // Replace with your desired AWS region
				},
			}
		}

		var s, err = session.NewSessionWithOptions(opts)

		if err != nil {
			fmt.Println("Something wrong when establishing the connection:", err)
		}
		var awsIamClient = goolsIam.IAMClient{
			SVC: iam.New(s),
		}

		u, err := awsIamClient.ListIamUsers()

		if err != nil {
			fmt.Println("Something wrong when generating the user list:", err)
		}

		dryRun, _ := cmd.Flags().GetBool("dry-run")

		awsIamClient.AccessKeyCleanup(dryRun, u, 90)
	},
}

func init() {
	iamCmd.AddCommand(keyCleanupCmd)

	keyCleanupCmd.Flags().BoolP("dry-run", "d", false, "don't actually deactivate the key, just show what would be done")
	keyCleanupCmd.Flags().Int32P("concurrency", "c", 1, "number of concurrent threads to use")
	keyCleanupCmd.Flags().Int32P("max-age", "m", 90, "maximum age of the key in days")

}
