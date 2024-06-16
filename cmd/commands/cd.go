package commands

import (
	"fmt"
	"os"
)

type CdCmd struct {
	Cmd
}

func (_ CdCmd) Run(args []string) (CommandResulter, error) {
	if len(args) == 0 {
		msg := "not enough parameters"
		return &CmdResult{Msg: msg, Code: 1}, fmt.Errorf(msg)
	}
	if len(args) > 1 {
		msg := "too many parameters"
		return &CmdResult{Msg: msg, Code: 1}, fmt.Errorf(msg)
	}
	err := os.Chdir(args[0])
	if err != nil {
		return &CmdResult{Msg: err.Error(), Code: 1}, err
	}
	return &CmdResult{}, nil
}

func (_ CdCmd) Description() string {
	return "cd is a shell builtin"
}
