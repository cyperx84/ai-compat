package main

import (
	"fmt"
	"os"

	"github.com/cyperx/ai-compat/internal/cli"
)

func main() {
	root := cli.NewRootCommand()
	if err := root.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
