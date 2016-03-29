package main

import (
	"os"
	"fmt"
	"github.com/kucuny/goenv/cli"
)

const (
	CliCommandHelp string = "help"
	CliCommandInstall string = "install"
	CliCommandCreateEnv string = "env"
)

func parseCommand() {
	if len(os.Args) == 0 {
		fmt.Println("Show help messages(no argument)")
		return
	}

	switch os.Args[1] {
	case CliCommandHelp:
		// Show help messages
		fmt.Println("Show help message")
	case CliCommandInstall:
		// Parse arguments
		// Install
		fmt.Println("Install go package")
		cli.Run()
	case CliCommandCreateEnv:
		// Parse arguments
		// Create env
		fmt.Println("Create virtual env")
	default:
		fmt.Println("Invalid command")
	}
}
