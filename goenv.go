package main

import (
	"fmt"
	"os"
	"github.com/kucuny/goenv/goinstall"
)

func getGoEnv() (string, string) {
	goRoot := os.Getenv("GOROOT")
	goPath := os.Getenv("GOPATH")

	return goRoot, goPath
}

func main() {
	goroot, gopath := getGoEnv()
	fmt.Println(goroot, gopath)
	pwd, err := os.Getwd()

	dirName := "test_env"

	err = os.MkdirAll(fmt.Sprintf(pwd+"/%s", dirName), 0755)

	if err != nil {
		fmt.Println("Cannot create directory")
	}

	gov := goinstall.NewGoVersion()

	gov.GetGoVersionListFromGolangOrg()
}
