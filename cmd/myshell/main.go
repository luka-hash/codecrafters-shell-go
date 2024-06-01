package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const prompt = "$ "

func main() {
loop:
	for {
		fmt.Fprintf(os.Stdout, "%s", prompt)
		line, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}
		switch args[0] {
		default:
			fmt.Fprintf(os.Stdout, "%s: command not found\n", args[0])
			break loop
		}

	}

}
