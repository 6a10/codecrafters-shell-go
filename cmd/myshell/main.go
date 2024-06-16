package main

import (
	"bufio"
	"fmt"
	cmd "github.com/codecrafters-io/shell-starter-go/cmd/commands"
	"os"
	"os/exec"
	"strings"
)

var NotImplementedErr = fmt.Errorf("not implemented")

var commands = map[string]cmd.Command{"echo": cmd.EchoCmd{}, "cat": cmd.CatCmd{}, "exit": cmd.ExitCmd{}, "pwd": cmd.PwdCmd{}}

func init() {
	commands["type"] = &cmd.TypeCmd{CmdMap: commands}
}

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
		if shellCmd, ok := commands[words[0]]; ok {
			res, err := shellCmd.Run(words[1:])
			if err != nil {
			} else {
			}
			fmt.Fprintf(os.Stdout, "%s\n", res.String())

		} else {
			execPath := cmd.FindExec(words[0])
			if execPath != "" {
				c := exec.Command(execPath, words[1:]...)
				c.Stdout = os.Stdout
				c.Stderr = os.Stderr
				err := c.Run()
				if err != nil {
					fmt.Fprintf(os.Stdout, err.Error())
				}
			} else {
				fmt.Fprintf(os.Stdout, "%s: command not found\n", words[0])
			}
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
