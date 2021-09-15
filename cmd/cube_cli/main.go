package main

import (
	"flag"
	"log"

	"github.com/DmiAS/cube_cli/internal/app/client"
	"github.com/DmiAS/cube_cli/internal/app/delivery/iproto"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 4 {
		log.Fatal("invalid number of args")
	}
	host, port := args[0], args[1]
	token, scope := args[2], args[3]

	cli, _ := iproto.NewClient(host, port)
	cube := client.NewCubeClient(cli)

	resp, err := cube.Send(token, scope)
	if err != nil {
		log.Fatal(err)
	}
	cube.PrintResponse(resp)
}
