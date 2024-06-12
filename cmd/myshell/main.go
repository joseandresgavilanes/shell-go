package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

// List of shell builtins
var builtins = map[string]bool{
	"echo": true,
	"exit": true,
	"type": true,
	"pwd": true,
	"cd": true,

}

// Function to check if a command is executable in any of the directories listed in PATH
func findExecutable(command string) (string, bool) {
	paths := strings.Split(os.Getenv("PATH"), ":")
	for _, dir := range paths {
		fullPath := filepath.Join(dir, command)
		if _, err := os.Stat(fullPath); err == nil {
			return fullPath, true
		}
	}
	return "", false
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
		case "pwd":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Fprintln(os.Stderr, "pwd: error retrieving current directory")
				continue
			}
			fmt.Fprintln(os.Stdout, dir)
		case "cd":
			if len(commands) < 2 {
				fmt.Fprintln(os.Stderr, "cd: usage: cd directory")
				continue
			}
			err := os.Chdir(commands[1])
			if err != nil {
				fmt.Fprintf(os.Stderr, "cd: %s: No such file or directory\n", commands[1])
				continue
			}
		case "type":
			if len(commands) < 2 {
				fmt.Fprintln(os.Stderr, "type: usage: type command")
				continue
			}
			for _, cmd := range commands[1:] {
				if builtins[cmd] {
					fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", cmd)
				} else if path, found := findExecutable(cmd); found {
					fmt.Fprintf(os.Stdout, "%s is %s\n", cmd, path)
				} else {
					fmt.Fprintf(os.Stdout, "%s: not found\n", cmd)
				}
			}
		default:
			// If not a builtin command, attempt to execute it as an external command
			if path, found := findExecutable(commands[0]); found {
				cmd := exec.Command(path, commands[1:]...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				if err := cmd.Run(); err != nil {
					fmt.Fprintf(os.Stderr, "%s: %v\n", commands[0], err)
				}
			} else {
				fmt.Fprintf(os.Stdout, "%s: command not found\n", message)
			}
		}
	}
}
