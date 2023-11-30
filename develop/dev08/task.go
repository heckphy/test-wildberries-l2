package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		prompt()

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		err = processInput(input)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func prompt() {
	fmt.Fprintln(os.Stdout)
	pwd()
	fmt.Fprint(os.Stdout, "$ ")
}

func processInput(input string) error {
	input = strings.TrimRight(input, "\r\n")

	if strings.Contains(input, "|") {
		return runPipe(input)
	}

	return runCommand(input)
}

func parseInput(input string) (string, []string) {
	args := strings.Fields(input)
	return args[0], args[1:]
}

func runCommand(input string) error {
	cmdName, cmdArgs := parseInput(input)

	switch cmdName {
	case "cd":
		if len(cmdArgs) == 0 {
			return chdir("")
		}

		return chdir(cmdArgs[0])
	case "pwd":
		return pwd()
	case "echo":
		return echo(cmdArgs...)
	case "exec":
		return exc(cmdName, cmdArgs...)
	case "exit":
		exit()
	}

	cmd := exec.Command(cmdName, cmdArgs...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func runPipe(input string) error {
	var b bytes.Buffer
	inputs := strings.Split(input, "|")

	for _, input := range inputs {
		input := strings.TrimSpace(input)
		cmdName, cmdArgs := parseInput(input)
		cmd := exec.Command(cmdName, cmdArgs...)

		cmd.Stdin = bytes.NewReader(b.Bytes())
		b.Reset()
		cmd.Stdout = &b

		err := cmd.Run()
		if err != nil {
			return err
		}
	}

	fmt.Fprint(os.Stdout, b.String())

	return nil
}

func chdir(filename string) error {
	if filename != "" {
		return os.Chdir(filename)
	}

	dir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	return os.Chdir(dir)
}

func pwd() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	fmt.Fprintln(os.Stdout, wd)

	return nil
}

func echo(args ...string) error {
	line := strings.Join(args, " ")
	_, err := fmt.Fprintln(os.Stdout, line)

	return err
}

func exc(cmdName string, cmdArgs ...string) error {
	binary, err := exec.LookPath(cmdName)
	if err != nil {
		return err
	}

	env := os.Environ()

	return syscall.Exec(binary, cmdArgs, env)
}

func exit() {
	os.Exit(0)
}
