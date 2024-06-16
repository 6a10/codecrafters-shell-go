package commands

import (
	_ "fmt"
	"os"
	"path/filepath"
)

type PwdCmd struct {
	Cmd
}

func (p PwdCmd) Run(args []string) (CommandResulter, error) {
	rpath, err := os.Getwd()
	if err != nil {
		return &CmdResult{Msg: err.Error(), Code: 1}, err
	}
	return &CmdResult{Msg: rpath}, nil
}

func (p PwdCmd) Description() string {
	return "pwd is a shell builtin"
}

type PWD struct {
	path string
}

func (p *PWD) Move(s string) error {
	newpath := filepath.Join(p.path, s)
	_, err := os.Stat(newpath)
	if err != nil {
		// if os.IsNotExist(err) {}
		return err
	}

	return nil
}
