package gosh

import (
	"github.com/nj-jay/go-shell/gosh/command"
	"io"
)


func interpret() error {

	// read user input
	formatInput, err := shellPrompt.formatInput()

	if err != nil {

		if err == io.EOF {

			return err
		}

		return err
	}


	// if no input is given, skip the cycle
	if len(formatInput) == 0 {

		return nil

	} else if formatInput[0] == "ls" {

		if len(formatInput) == 1 {

			command.Ls(".")

		} else {

			command.Ls(formatInput[1])
		}


	} else if formatInput[0] == "cd" {
		    command.Cd(formatInput)
		}


	return nil
}



