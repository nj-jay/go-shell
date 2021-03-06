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

	} else if IsBuilt(argv[0]) {

		//显得更美观
		fmt.Println()

		//将敲击的命令写进.go-shell.history

		if argv[0] != "history" {

			command.WriteHistory(argv)

		}

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

		case "cat":

			command.Cat(argv)

		case "gedit":

			command.Gedit(argv)

		case "history":

			command.History(argv)

		case "rm":

			command.Rm(argv)

		case "mv":

			command.Mv(argv)

		case "update":

			command.Update(argv)

		//case "hello":
		//
		//	//后续转为这种开发
		//	//更为方便哈哈哈哈
		//	//真牛逼噢
		//	//稍微更难配置一点
		//	//但是缺更能增加用户的体验不是嘛
		//	cli(argv)

		case "pwd":


			cli(argv)


		}

	} else {

		Builtin(argv)

	}
}


