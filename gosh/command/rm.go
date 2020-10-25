package command

import (

	"fmt"
	"os"

)

const (

	ContextRm = `
	rm hello.c			删除当前目录下hello.c文件
	rm d:/test/hello.c		删除D:/test目录下的hello.c文件
	rm -rf dir			删除当前目录下的dir目录
	rm -rf d:/test/dir	删除d:/test目录下的dir目录
`

)
type HelpRm struct {


}

func (*HelpRm) help() string {

	return ContextRm
}

func Rm(argv []string) {

	//rm filename
	if len(argv) == 1 {

		fmt.Println("go-shell: rm: 缺少操作对象")

	} else if len(argv) == 2 && (argv[1] != "--help" && argv[1] != "-h") {

		err := os.Remove(argv[1])

		if err != nil {

			fmt.Println("go-shell: rm: 没有这个文件")

		}

	} else if len(argv) == 3 && argv[1] == "-rf" {

		err := os.RemoveAll(argv[2])

		if err != nil {

			fmt.Println("go-shell: rm: 没有这个目录")

		}

	} else if len(argv) == 2 && (argv[1] == "--help" || argv[1] == "-h"){

		helpRm := new(HelpRm)

		fmt.Println(helpRm.help())

	}
}