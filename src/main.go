package main

import (

	"github.com/nj-jay/go-shell/gosh"
	"os"
	"os/exec"
)


//萌生新的想法
//oh yes!!!
//可以更好的使用go-shell
//用户只用配置好环境变量
//非常的nicer!!!!!!!!


func main() {


	c := "set path=%path%;D:/go-shell/bin"

	cmd := exec.Command("cmd.exe", "/c", c)

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr

	cmd.Run()

	gosh.Run()
}