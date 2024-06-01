package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const prompt = "$ "

func main() {
	builtins := []string{"echo", "exit", "type"}
	for {
		fmt.Fprintf(os.Stdout, "%s", prompt)
		line, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}
		command, args, _ := next(args)
		switch command {
		case "echo":
			rest, _ := strings.CutPrefix(line, command)
			fmt.Fprintf(os.Stdout, "%s\n", strings.TrimSpace(rest))
		case "type":
			cmd, _, err := next(args)
			if err != nil {
				fmt.Fprintf(os.Stdout, "%s\n", "type builtin requires 1 argument")
			}
			if slices.Contains(builtins, cmd) {
				fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", cmd)
			} else {
				fmt.Fprintf(os.Stdout, "%s not found\n", cmd)
			}
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
