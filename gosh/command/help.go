package command

import "fmt"

//Helper接口
type Helper interface {

	Help()

}

const (

	ContextHelp = `
go-shell program v0.0.1

Usage:
	command [option]

Available Commands:

	ls			列出指定目录下的文件
	cd			进入到指定的目录
	mkdir			指定位置创建目录
	touch			指定位置创建文件
	help			查看支持的命令
	cls			清屏

Option:

	command --help			可以获得更多关于命令的使用

	for example, You can use "cd --help" to get more formation!
`
)

type HelpCommand struct {

}

func (*HelpCommand) help() string {

	return ContextHelp


}

func Help() {

	help := new(HelpCommand)

	helpText := help.help()

	fmt.Println(helpText)

}