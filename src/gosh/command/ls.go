package command

import (

	"bufio"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"os"

)

const (

	ContextLs = `
	ls [option]

	ls			列出当前目录所有的文件
	ls ..			列出上一级目录所有的文件
	ls ../test		列出上一级目录的test目录下的所有文件
	ls work			列出当前目录下的work目录下所有的文件
	ls d:			列出d盘下的所有文件
`

)
type HelpLs struct {


}

func (*HelpLs) help() string {

	return ContextLs
}


type UnixLikeLs struct {

	stdout string

	out string

}

var (

	ls UnixLikeLs

	//output string

)

// `>`  or `>>`
func (*UnixLikeLs) toFile(name, content, flag string) {

	if flag == ">" {

		//将string写进file中 覆盖写入 必须加 os.O_TRUNC 这是个坑啊!!!!!
		f, err := os.OpenFile(name, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0666)

		defer f.Close()

		if err != nil {

			fmt.Println("文件打开失败")

		}

		outputWriter := bufio.NewWriter(f)

		_, _ = outputWriter.WriteString(content + "\n")

		_ = outputWriter.Flush()

	} else if flag == ">>" {

		f, err := os.OpenFile(name, os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0666)

		defer f.Close()

		if err != nil {

			fmt.Println("文件打开失败")

		}

		outputWriter := bufio.NewWriter(f)

		_, _ = outputWriter.WriteString(content + "\n")

		_ = outputWriter.Flush()

	}
}

func Ls(argv []string) {

	if len(argv) == 1 {

		argv = []string{"ls", "."}

	}

	if argv[1] == "--help" || argv[1] == "-h" {

		helpLs := new(HelpLs)

		fmt.Println(helpLs.help())

	} else if len(argv) == 3 && (argv[1] == ">" || argv[1] == ">>"){

		output := ls.out

		//Touch([]string{"touch", argv[2]})

		//Rm([]string{"rm", argv[2]})

		//判断是">" 还是">>"
		flag := argv[1]

		ls.toFile(argv[2], output, flag)

	} else {

		fmt.Println(ls.output(argv[1][:]))

		}
	}

func (l *UnixLikeLs) output(dir string) string {

	l.stdout = ""

	info, err := ioutil.ReadDir(dir)

	if err != nil {

		fmt.Println("no this dictionary")
	}

	for _, f := range info {

		if ok := f.IsDir(); ok {


			blue := color.New(color.FgBlue)
			//l.stdout += color.BlueString(f.Name()) + "\n"
			l.stdout += blue.Sprintf("%s", f.Name()) + "\n"

		} else {

			l.stdout += f.Name() + "\n"

		}
	}

	l.out = l.stdout

	return l.stdout
}