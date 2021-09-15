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

	cli, err := iproto.NewClient(host, port)
	if err != nil {
		log.Fatal("can't connect to cube service due to", err)
	}
	cube := client.NewCubeClient(cli)

	resp, err := cube.Send(token, scope)
	if err != nil {
		log.Fatal("internal error = ", err)
	}
	cube.PrintResponse(resp)
}
