package main

import (
	"os"

	"github.com/shihanng/devto/cmd"
)

func main() {
	c, sync := cmd.New(os.Stdout)
	defer sync()

	if err := c.Execute(); err != nil {
		os.Exit(1)
	}
}
