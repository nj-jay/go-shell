package gosh

import (

	"os"
	"os/exec"

)


/*

设置环境变量等，支持用户自由设置

 */


func cli(argv []string) {

	var command  = "d:/go-shell/bin/"

	for _, value := range argv {

		command += value + " "

	}

	cmd := exec.Command("cmd.exe", "/c", command)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	cmd.Run()
}