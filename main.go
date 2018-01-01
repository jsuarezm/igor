package main

import "github.com/jsuarezm/igor/cmd"

var (
	VERSION = "0.0.1"
)

func main() {
	cmd.Execute(VERSION)
}
