package gosh

import (

	"fmt"
	"github.com/nj-jay/go-shell/gosh/command"

)

//判断用户输入的命令

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

		case "help":

			command.Help()

		case "wget":

			command.Wget(argv)

		case "exit":

			command.Exit(argv)

		case "cls":

			command.Cls(argv)

		default:

			fmt.Println("go-shell: command not found")

			fmt.Println()
		}
	}

}


