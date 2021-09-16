package main

import (
	"flag"
	"os"

	"github.com/DmiAS/cube_cli/command"
)

func main() {
	flag.Parse()
	args := flag.Args()
	os.Exit(command.Run(args))
}
