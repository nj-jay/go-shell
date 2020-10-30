package gosh

import (

	"fmt"
	"os"
	"os/exec"

)

//2020.10.25新的想法
//支持对windows其他命令的支持
//如python java等
//大大提升了go-shell的可ls用性

func Builtin(argv []string) {

	var formatInput string

	//formatInput = argv[0] + " "

	for _, v := range argv {

		formatInput += v + " "
	}

	fmt.Println()

	cmd := exec.Command("cmd.exe", "/c", formatInput)

	cmd.Stdin = os.Stdin

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr

	cmd.Run()

	//if err != nil {
	//
	//	fmt.Println("go-shell: command not found")
	//
	//}

	}

func IsBuilt(argv string) bool{

	switch argv {

	case "ls", "cd", "mkdir", "touch", "cat", "rm", "wget",
	"history", "exit", "gedit", "help", "mv", "update":

		return true

	default:

		return false

	}
}