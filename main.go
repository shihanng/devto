package main

import (
	"os"

	"github.com/shihanng/devto/cmd"
)

func main() {
	c, sync := cmd.New(os.Stdout)
	_ = c.Execute()

	sync()
}
