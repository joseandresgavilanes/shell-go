package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

		fmt.Fprint(os.Stdout, "$ ")
	
	command, _ := reader.ReadString('\n')
	command = strings.TrimSuffix(command, "\n")
	fmt.Printf("%s: command not found\n", strings.TrimRight(command, "\n"))
	
	cmd := strings.TrimSpace(command);

	if cmd == "exit 0" {
		
		os.Exit(0)
	}
}
