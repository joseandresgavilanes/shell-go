package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// List of shell builtins
var builtins = map[string]bool{
	"echo": true,
	"exit": true,
	"type": true,
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Fprint(os.Stdout, "$ ")
		message, err := reader.ReadString('\n')
		if err != nil {
			os.Exit(1)
		}
		message = strings.TrimSpace(message)
		commands := strings.Split(message, " ")

		switch commands[0] {
		case "exit":
			if len(commands) > 1 {
				code, err := strconv.Atoi(commands[1])
				if err != nil {
					fmt.Fprintln(os.Stderr, "exit: invalid exit status")
					continue
				}
				os.Exit(code)
			}
			os.Exit(0)
		case "echo":
			fmt.Fprintln(os.Stdout, strings.Join(commands[1:], " "))
		case "type":
			if len(commands) < 2 {
				fmt.Fprintln(os.Stderr, "type: usage: type command")
				continue
			}
			for _, cmd := range commands[1:] {
				if builtins[cmd] {
					fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", cmd)
				} else {
					fmt.Fprintf(os.Stdout, "%s: not found\n", cmd)
				}
			}
		default:
			fmt.Fprintf(os.Stdout, "%s: command not found\n", message)
		}
	}
}
