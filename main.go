// Package main contains a small demo program showing how to use a terminalRunner
// to parse a YarnSpinner dialogue file, then run and display it in a terminal.
package main

import (
	"fmt"
	"os"

	"github.com/RemiEven/ysgo/terminal"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("missing file path argument")
		os.Exit(1)
	}

	terminalRunner, err := terminal.NewRunner(os.Args[1], "")
	if err != nil {
		panic(fmt.Errorf("failed to create terminal runner: %w", err))
	}

	if err := terminalRunner.Run(); err != nil {
		panic(fmt.Errorf("terminal runner failed to run: %w", err))
	}
}
