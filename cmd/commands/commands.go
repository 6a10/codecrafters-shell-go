package commands

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var NotImplementedErr error = fmt.Errorf("not Implemented")

type CommandResulter interface {
	String() string
	Value() uint8
}

type Command interface {
	Run([]string) (CommandResulter, error)
	SanitizeString(string) []string
	Description() string
}

type CmdResult struct {
	Msg  string
	Code uint8
}

func (c CmdResult) String() string {
	return c.Msg
}

func (c CmdResult) Value() uint8 {
	return c.Code
}

type Cmd struct{}

func (c Cmd) Run(_ []string) (CommandResulter, error) {
	return nil, NotImplementedErr
}

func (c Cmd) SanitizeString(s string) []string {
	arr := strings.Split(strings.ReplaceAll(s, "\n", ""), " ")
	result := make([]string, 0, len(arr))
	for _, w := range arr {
		if w != "" {
			result = append(result, w)
		}
	}
	return result
}

func (c Cmd) Description() string {
	return ""
}

type ExitCmd struct {
	Cmd
	defaultExitCode int
}

func (ec ExitCmd) Run(args []string) (CommandResulter, error) {
	var exitCode = ec.defaultExitCode
	if len(args) != 0 {
		i, err := strconv.Atoi(args[0])
		if err != nil {
			return &CmdResult{Msg: fmt.Sprintf("parsing \"%s\": invalid syntax", args[0]), Code: 1}, fmt.Errorf("ExitCmd: %w", err)
		}
		exitCode = i
	}
	os.Exit(exitCode)
	return &CmdResult{}, nil
}

func (ec ExitCmd) Description() string {
	return "exit is a shell builtin"
}

type EchoCmd struct {
	Cmd
}

func (ec EchoCmd) Run(args []string) (CommandResulter, error) {
	s := strings.Join(args, " ")
	return &CmdResult{Msg: s, Code: 0}, nil
}

func (ec EchoCmd) Description() string {
	return "echo is a shell builtin"
}

type CatCmd struct {
	Cmd
}

func (cc CatCmd) Run(args []string) (CommandResulter, error) {
	return &CmdResult{Code: 1}, NotImplementedErr
}

func (cc CatCmd) Description() string {
	return "cat is /bin/cat"
}

type TypeCmd struct {
	Cmd
	CmdMap map[string]Command
}

func (tc *TypeCmd) Run(args []string) (CommandResulter, error) {
	if len(args) == 0 {
		return &CmdResult{Msg: "not enough arguments", Code: 1}, nil
	}
	if len(args) > 1 {
		return &CmdResult{Msg: "to many arguments", Code: 1}, nil
	}
	if cmd, ok := tc.CmdMap[args[0]]; ok {
		return &CmdResult{Msg: cmd.Description(), Code: 1}, nil
	} else {
		pathRes := FindExec(args[0])
		if pathRes != "" {
			return &CmdResult{Msg: fmt.Sprintf("%s is %s", args[0], pathRes)}, nil
		}
		// return &CmdResult{Msg: fmt.Sprintf("%s: not found", args[0]), Code: 1}, nil
	}
	return &CmdResult{Msg: fmt.Sprintf("%s: not found", args[0]), Code: 1}, nil
	// return &CmdResult{Msg: "not implemented", Code: 1}, NotImplementedErr
}

func (tc TypeCmd) Description() string {
	return "type is a shell builtin"
}
