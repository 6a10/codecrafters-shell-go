package commands

import (
	"fmt"
	"os"
	"strings"
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
	home := os.Getenv("HOME")
	s := strings.ReplaceAll(args[0], "~", home)
	err := os.Chdir(s)
	if err != nil {
		if os.IsNotExist(err) {
			return &CmdResult{Msg: fmt.Sprintf("%s: No such file or directory", args[0]), Code: 1}, err
		}
		return &CmdResult{Msg: err.Error(), Code: 1}, err
	}
	return &CmdResult{}, nil
}

func (_ CdCmd) Description() string {
	return "cd is a shell builtin"
}
