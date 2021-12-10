package main

import (
	"os"

	"github.com/ess/kit/cmd/kit/commands"
)

func main() {
	if commands.Execute() != nil {
		os.Exit(1)
	}
}
