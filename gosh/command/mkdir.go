package command

import (

	"fmt"
	"os"

)

const (

	ContextMkdir = `
	mkdir directoryNamePath(指定path)

	mkdir directoryName			在当前目录下创建目录
	mkdir ../directoryName			在上一级目录下创建目录
	mkdir /path/to/directoryName		在/path/to目录下创建目录
`

)
type HelpMkdir struct {


}

func (*HelpMkdir) help() string {

	return ContextMkdir
}

func Mkdir(argv []string) {


	if argv[1] == "--help" || argv[1] == "-h" {

		helpMkdir := new(HelpMkdir)

		fmt.Println(helpMkdir.help())

	} else {

		//判断目录存不存在
		b := PathExists(argv[1])

		//if not exist
		if !b {

			err := os.Mkdir(argv[1], 0777)

			if err != nil {

				fmt.Println("mkdir: 无法创建目录: 权限不够")

			}
		}

	}

}

func PathExists(path string) bool {

	_, err := os.Stat(path)

	if err == nil {

		return true
	}
	if os.IsNotExist(err) {

		return false
	}

	return false
}