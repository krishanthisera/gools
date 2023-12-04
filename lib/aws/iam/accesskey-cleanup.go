package iam

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
)

func (C IAMClient) AccessKeyCleanup(dryRun bool, iamUsers []*iam.User, expiry int) error {

	// Iterate through users
	for _, u := range iamUsers {

		// List access keys for each user
		keys, err := C.SVC.ListAccessKeys(&iam.ListAccessKeysInput{
			UserName: u.UserName,
		})
		if err != nil {
			fmt.Println("error listing access keys for user", *u.UserName, ":", err)
			continue
		}

		// Iterate through access keys
		for _, key := range keys.AccessKeyMetadata {

			// Get last used information for the access key
			lastUsedOutput, err := C.SVC.GetAccessKeyLastUsed(&iam.GetAccessKeyLastUsedInput{
				AccessKeyId: key.AccessKeyId,
			})

			if err != nil {
				fmt.Println("error getting last used information for access key", *key.AccessKeyId, ":", err)
				continue
			}

			// Calculate the duration since last usage
			var inactiveDays int
			if lastUsedOutput.AccessKeyLastUsed != nil && lastUsedOutput.AccessKeyLastUsed.LastUsedDate != nil {
				lastUsedTime := *lastUsedOutput.AccessKeyLastUsed.LastUsedDate
				inactiveDays = int(time.Since(lastUsedTime).Hours() / 24)
			} else {
				fmt.Println("access key", *key.AccessKeyId, "for user", *u.UserName, "never used.")
				continue
			}

			// Check if the access key has been inactive for more than 6 months (180 days)
			if inactiveDays > expiry {
				fmt.Println("access key", *key.AccessKeyId, "for user", *u.UserName, "inactive for more than ", expiry, "days.")
				if !dryRun {
					fmt.Println("deactivating access key", *key.AccessKeyId, "for user", *u.UserName)
					_, err := C.SVC.UpdateAccessKey(&iam.UpdateAccessKeyInput{
						AccessKeyId: key.AccessKeyId,
						Status:      aws.String("Inactive"),
						UserName:    u.UserName,
					})
					if err != nil {
						fmt.Println("error deactivating access key", *key.AccessKeyId, "for user", *u.UserName, ":", err)
					}
				}
			}
		}
	}
	return nil
}
