package main

import "github.com/shihanng/devto/cmd"

func main() {
	c, sync := cmd.New()
	c.Execute()
	sync()
}
