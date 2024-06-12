package main

import (
	"bufio"
	"fmt"
	"os"
	_ "os/signal"
	"strings"
	_ "syscall"
)

type Command interface{}

type Cmd struct{}

var commands = map[string]Command{"echo": Cmd{}, "cd": Cmd{}}

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

		words := strings.Split(cmdString[:len(cmdString)-1], " ")
		if len(words) == 0 {
			continue
		}
		if _, ok := commands[words[0]]; ok {
		} else {
			fmt.Fprintf(os.Stdout, "%s: command not found\n", words[0])
			os.Exit(1)
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
