/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"krishanthisera/gools/cmd"
	_ "krishanthisera/gools/cmd/aws"
	_ "krishanthisera/gools/cmd/prerender"
)

func main() {
	cmd.Execute()
}
