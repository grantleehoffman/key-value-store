package main

import (
	"fmt"
	"github.com/grantleehoffman/key-value-store/cli/cmd"
	"os"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
