package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path"
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
			arg, _ := peek(args)
			if arg[0] == '$' {
				v, _ := os.LookupEnv(arg[1:])
				fmt.Fprintf(os.Stdout, "%s\n", v)
			} else {
				rest, _ := strings.CutPrefix(line, command)
				rest = strings.TrimSpace(rest)
				fmt.Fprintf(os.Stdout, "%s\n", rest)
			}
		case "type":
			cmd, _, err := next(args)
			if err != nil {
				fmt.Fprintf(os.Stdout, "%s\n", "type builtin requires 1 argument")
			}
			if slices.Contains(builtins, cmd) {
				fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", cmd)
			} else if e, ok := searchPath(cmd); ok {
				fmt.Fprintf(os.Stdout, "%s is %s\n", cmd, e)
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

func searchPath(cmd string) (string, bool) {
	PATH, _ := os.LookupEnv("PATH")
	for _, subpath := range strings.Split(PATH, ":") {
		files, _ := os.ReadDir(subpath)
		for _, item := range files {
			if !item.IsDir() && item.Name() == cmd {
				return path.Join(subpath, item.Name()), true
			}
		}
	}
	return "", false
}
