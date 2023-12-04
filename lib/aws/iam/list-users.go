package iam

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/iam"
)

func (C IAMClient) ListIamUsers() ([]*iam.User, error) {

	// List IAM users
	listUsersInput := &iam.ListUsersInput{}
	listUsersOutput, err := C.SVC.ListUsers(listUsersInput)
	if err != nil {
		return listUsersOutput.Users, fmt.Errorf("error listing users: %v", err)
	}

	return listUsersOutput.Users, nil
}
