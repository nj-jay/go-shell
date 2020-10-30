package gosh

import (

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

	shellCommand(formatInput)

	return nil
}