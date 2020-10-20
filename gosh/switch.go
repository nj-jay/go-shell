package gosh

import (
	"fmt"
	"github.com/nj-jay/go-shell/gosh/command"
)

func shellCommand(argv []string) {

	//如果直接敲空格　cd到当前目录
	if len(argv) == 0 {

		command.Cd([]string{"cd", Usr.CurrentPath})

	} else {

		switch argv[0] {

		case "cd":

			command.Cd(argv)

		case "ls":

			command.Ls(argv)

		case "mkdir":

			command.Mkdir(argv)

		case "touch":

			command.Touch(argv)

		default:

			fmt.Println("go-shell: command not found")

			fmt.Println()
		}
	}



}