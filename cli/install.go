package cli

import (
	"flag"
	"fmt"
	"os"
)

var (
	FlagSetInstall = flag.NewFlagSet("install", flag.ContinueOnError)
	FlagInstallVersion = FlagSetInstall.String("v", "latest", "Golang install target version")
)

func init() {
	FlagSetInstall.Parse(os.Args[2:])
}

func Run() {
	fmt.Println(*FlagInstallVersion)
}