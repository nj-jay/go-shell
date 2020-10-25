package gosh

import (

	"fmt"
	"os"
	"os/exec"
)

//2020.10.25新的想法
//支持对windows其他命令的支持
//如python java等
//大大提升了go-shell的可用性

func Builtin(argv []string) string {

	var formatInput string

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

	return formatInput
}

func IsBuilt(argv string) bool{

	switch argv {

	case "ls", "cd", "mkdir", "touch", "cat", "rm", "wget", "history", "exit", "gedit", "help":

		return true

	default:

		return false

	}
}