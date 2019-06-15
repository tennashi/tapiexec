package tapiexec

import (
	"encoding/json"
	"fmt"

	"golang.org/x/crypto/ssh/terminal"
)

type Call struct {
	FuncName string
	Args     []string
}

type Drop struct {
	FileName string
	Options  map[string]string
}

type Tapi interface {
	Run() error
}

func CallAPI(funcName string, args []string) Tapi {
	return &Call{
		FuncName: funcName,
		Args:     args,
	}
}

func DropAPI(fileName string, options map[string]string) Tapi {
	return &Drop{
		FileName: fileName,
		Options:  options,
	}
}

func (c *Call) Run() error {
	v := []interface{}{"call", c.FuncName, c.Args}
	jsonByte, err := json.Marshal(v)
	if err != nil {
		return err
	}
	call := "\x1b]51;" + string(jsonByte) + "\x07"
	fmt.Printf(call)

	return nil
}

func (d *Drop) Run() error {
	v := []interface{}{"drop", d.FileName, d.Options}
	jsonByte, err := json.Marshal(v)
	if err != nil {
		return err
	}
	call := "\x1b]51;" + string(jsonByte) + "\x07"
	fmt.Printf(call)

	return nil
}

func WaitMsg(msg string) error {
	for {
		ret, err := terminal.ReadPassword(0)
		if err != nil {
			return err
		}
		if string(ret) == msg {
			break
		}
	}
	return nil
}
