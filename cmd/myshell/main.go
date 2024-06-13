package main

import (
	"bufio"
	"fmt"
	"os"
	_ "os/signal"
	"strconv"
	"strings"
	_ "syscall"
)

var NotImplementedErr = fmt.Errorf("not implemented")

type CommandResulter interface {
	String() string
	Value() uint8
}

type Command interface {
	Run([]string) (CommandResulter, error)
	SanitizeString(string) []string
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

type EchoCmd struct {
	Cmd
}

func (ec EchoCmd) Run(args []string) (CommandResulter, error) {
	s := strings.Join(args, " ")
	return &CmdResult{Msg: s, Code: 0}, nil
}

var commands = map[string]Command{"echo": EchoCmd{}, "cd": Cmd{}, "exit": ExitCmd{}}

// var controls = map[string]error{"^C": fmt.Errorf("exit command"), "^D": fmt.Errorf("close command")}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

	reader := bufio.NewReader(os.Stdin)

	for {
		// Uncomment this block to pass the first stage
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprint(os.Stderr, "error reading command")
			os.Exit(1)
		}
		cmdString = cmdString[:len(cmdString)-1]

		splited := strings.Split(cmdString, " ")
		words := make([]string, 0, len(splited))
		for _, w := range splited {
			if w != "" {
				words = append(words, w)
			}
		}
		if len(words) == 0 {
			continue
		}
		if cmd, ok := commands[words[0]]; ok {
			res, err := cmd.Run(words[1:])
			if err != nil {
			} else {
			}
			fmt.Fprintf(os.Stdout, "%s\n", res.String())

		} else {
			fmt.Fprintf(os.Stdout, "%s: command not found\n", words[0])
			// os.Exit(1)
		}
	}
}

// func init() {
// 	c := make(chan os.Signal)
// 	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
// 	go func() {
// 		<-c
// 		// Run Cleanup
// 		os.Exit(1)
// 	}()
// }
