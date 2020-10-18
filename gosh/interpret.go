package gosh

import (

	"fmt"
	"github.com/nj-jay/go-shell/gosh/command"
	"io"
	"strings"
)

func interpret() error {

	// read user input
	input, err := readInput()

	if err != nil {

		if err == io.EOF {

			return err
		}

		return err
	}

	input = strings.TrimRight(input, "\r\n")

	fmt.Print("$ ")

	// if no input is given, skip the cycle
	if input == "" {

		return nil

	} else if input == "ls" {

		command.Ls(".")

	} else {

		fmt.Println("go-shell: command not found")

		fmt.Println()
	}


	return nil
}

func readInput() (string, error) {

	return reader.ReadString('\n')
}