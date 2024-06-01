package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const prompt = "$ "

func main() {
	for {
		fmt.Fprintf(os.Stdout, "%s", prompt)
		line, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}
		command, args, _ := next(args)
		switch command {
		case "exit":
			exit_code_str, _, err := next(args)
			if err != nil {
				os.Exit(0)
			}
			exit_code, err := strconv.Atoi(exit_code_str)
			if err != nil {
				panic(err)
			}
			os.Exit(exit_code)
		default:
			fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
		}
	}
}

func next(ss []string) (string, []string, error) {
	if len(ss) == 0 {
		var z string
		return z, ss, errors.New("slice is empty")
	}
	return ss[0], ss[1:], nil
}

func peek(ss []string) (string, error) {
	if len(ss) == 0 {
		return "", errors.New("lice is empty")
	}
	return ss[0], nil
}
