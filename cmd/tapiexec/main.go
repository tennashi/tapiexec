package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/tennashi/tapiexec"
)

var buf = bufio.NewScanner(os.Stdin)

func main() {
	if len(os.Args) < 3 {
		return
	}

	switch os.Args[1] {
	case "call":
		err := handleCall(os.Args[2:])
		handleError(err)
	case "drop":
		err := handleDrop(os.Args[2:])
		handleError(err)
	}
}

func handleCall(args []string) error {
	if len(args) == 0 {
		return NewErrInvalidArgs()
	}

	funcName := args[0]
	if !strings.HasPrefix(funcName, "Tapi_") {
		funcName = "Tapi_" + funcName
	}

	tapiArgs := []string{}
	if len(args) > 1 {
		tapiArgs = args[1:]
	}

	tapi := tapiexec.CallAPI(funcName, tapiArgs)
	return tapi.Run()
}

func handleDrop(args []string) error {
	if len(args) == 0 {
		return NewErrInvalidArgs()
	}

	fileName := args[0]

	options := map[string]string{}
	if len(args) > 1 {
		for _, opt := range args[1:] {
			option := strings.SplitN(opt, ":", 2)
			if len(option) != 2 {
				continue
			}
			options[option[0]] = option[1]
		}
	}

	tapi := tapiexec.DropAPI(fileName, options)
	return tapi.Run()
}

type ErrInvalidArgs struct {
	msg string
}

func NewErrInvalidArgs() ErrInvalidArgs {
	return ErrInvalidArgs{"invalid arguments"}
}

func (e ErrInvalidArgs) Error() string {
	return e.msg
}

func handleError(err error) {}
