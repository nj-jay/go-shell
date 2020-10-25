package command

import (
	"fmt"
	"os"
)

const (

	ContextMv = `
	mv path path
	mv hello.c main.c			将当前目录下的hello.c改为main.c
	mv hello.c ../hello.c			将当前目录下的hello.c移动到上一级目录并命名为hello.c
	mv hello.c ../world.c			将当前目录下的hello.c移动到上一级目录并命名为world.c
	mv d:/test/hello.c ./hello.c		将d:/test/hello.c移动到当前目录
	mv d:/test/hello.c e:/test/hello.c	将d:/test目录下的hello.c移动到e:/test
`
)

type HelpMv struct {

}

func (*HelpMv) help() string {

	return ContextMv

}

func Mv(argv []string) {

	//input == mv --help or mv -h
	if len(argv) == 2 && (argv[1] == "--help" || argv[1] == "-h") {

		helpMv := new(HelpMv)

		fmt.Println(helpMv.help())

	} else if len(argv) == 1 || len(argv) == 2 {

		//input == mv or mv hello.c
		fmt.Println("go-shell: mv: 缺少操作对象")

	} else if len(argv) == 3 {

		err := os.Rename(argv[1], argv[2])

		if err != nil {

			fmt.Println("go-shell: mv: 没有这个文件")

		}

	}

}

//判断是否是目录

func IsDir(name string) bool {

	fileInfo, err := os.Stat(name)

	if err != nil {

		return false

	}

	b := fileInfo.IsDir()

	return b

}