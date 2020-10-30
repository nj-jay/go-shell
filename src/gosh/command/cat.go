package command

import (
	"bufio"
	"fmt"
	"github.com/spf13/cast"
	"io"
	"io/ioutil"
	"os"
)

/*
查看文本文件的命令
 */


const (

	ContextCat = `
			cat hello.c			查看当前目录下文本文件hello.c的内容
			cat ../hello.c			查看上一级目录下文本文件hello.c的内容
			cat d:/test/hello.c		查看d:/test目录下hello.c的内容
			cat -n hello.c			显示行数
`

)
type HelpCat struct {


}

func (*HelpCat) help() string {

	return ContextCat

}


func Cat(argv []string) {

	if len(argv) == 1 {

		fmt.Println("go-shell: cat缺少操作对象")

	} else if argv[1] == "--h" || argv[1] == "-h" {

		helpCat := new(HelpCat)

		fmt.Println(helpCat.help())

	} else if len(argv) == 3 && argv[1] == "-n" {

		//cat -n test.txt
		//有一个bug尚未解决 当文件只有一行时且没有换行符的时候cat -n无法显示

		var line int = 1

		f, err := os.Open(argv[2])

		defer f.Close()

		if err != nil {

			fmt.Println("go-shell: cat -n 没有这个文件")

			return

		}

		reader := bufio.NewReader(f)

		for  {

			lineString, err := reader.ReadString('\n')

			if err == io.EOF {

				fmt.Println()

				return
			}

			fmt.Print(cast.ToString(line) + " " + lineString)

			line = line + 1 //行数加1

		}

	} else if len(argv) == 3 && (argv[1] == "-h" || argv[1] == "--help") {

			helpCat := new(HelpCat)

			fmt.Println(helpCat.help())

	} else {

		f, err := os.Open(argv[1])

		defer f.Close()

		if err != nil {

			fmt.Println("没有这个文件")

			//不执行后续操作
			return

		}

		contentByte, err := ioutil.ReadAll(f)

		if err != nil {

			fmt.Println("请打开文本文件")

		}

		fmt.Println(string(contentByte))

	}
}