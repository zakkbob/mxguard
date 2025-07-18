/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	rootcmd "github.com/zakkbob/mxguard/cmd"
	_ "github.com/zakkbob/mxguard/cmd/start"
	_ "github.com/zakkbob/mxguard/cmd/stop"
	_ "github.com/zakkbob/mxguard/cmd/user"
	_ "github.com/zakkbob/mxguard/cmd/user/alias"
)

func main() {
	rootcmd.Execute()
}
