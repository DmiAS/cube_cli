package command

import (
	"fmt"

	"github.com/DmiAS/cube_cli/internal/app/cli"
	"github.com/DmiAS/cube_cli/internal/app/delivery/iproto"
)

const (
	argsError = iota
	connectionError
	internalError
	argsLen = 4
)

func Run(args []string) int {
	if len(args) < argsLen {
		fmt.Println("not enough arguments for cli call")
		return argsError
	}
	host, port := args[0], args[1]
	token, scope := args[2], args[3]

	proto, err := iproto.NewClient(host, port)
	if err != nil {
		fmt.Println("can't connect to cube service due to", err)
		return connectionError
	}
	cube := cli.NewCubeClient(proto)

	resp, err := cube.Send(token, scope)
	if err != nil {
		fmt.Println("internal error = ", err)
		return internalError
	}
	cube.PrintResponse(resp)
	return 0
}
