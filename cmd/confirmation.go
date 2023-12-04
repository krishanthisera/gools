package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ConfirmAction() bool {
	fmt.Print("Are you sure you want to perform this action? (yes/no): ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)

	if input == "yes" {
		return true
	} else {
		return false
	}
}
