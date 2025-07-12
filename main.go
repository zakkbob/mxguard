/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/zakkbob/mxguard/cmd"
	_ "github.com/zakkbob/mxguard/cmd/start"
	_ "github.com/zakkbob/mxguard/cmd/stop"
	_ "github.com/zakkbob/mxguard/cmd/user"
)

func main() {
	cmd.Execute()
}
