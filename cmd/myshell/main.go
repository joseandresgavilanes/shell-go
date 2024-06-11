package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Fprint(os.Stdout, "$ ")
	for {
		cmd, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		cmd = strings.TrimSpace(cmd)
		switch cmd {
		case "exit 0":
			os.Exit(0)

		case "echo":
			fmt.Println(cmd[5:])
		}
	}
}
