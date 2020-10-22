package gosh

import (

	"io"

)

//shell解释器

func interpret() error {

	// read user input
	formatInput, err := shellPrompt.FormatInput()

	if err != nil {

		if err == io.EOF {

			return err
		}

		return err
	}

	shellCommand(formatInput)

	return nil
}



