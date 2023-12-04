package iam

import "github.com/aws/aws-sdk-go/service/iam"

type IAMGools interface {
	AccessKeyCleanup(dryRun bool, iamUsers *[]iam.User) error
	ListIamUsers() ([]*iam.User, error)
}

type IAMClient struct {
	SVC *iam.IAM
}
