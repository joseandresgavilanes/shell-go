package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Logs from your program will appear here!")

	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	bufio.NewReader(os.Stdin).ReadString('\n')
}
