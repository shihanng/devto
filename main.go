package main

import "github.com/shihanng/devto/cmd"

func main() {
	c, sync := cmd.New()
	_ = c.Execute()

	sync()
}
