package command

import (

	"fmt"
	"github.com/fatih/color"
	"io/ioutil"

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


func Ls(argv []string) {

	if len(argv) == 1 {

		argv = []string{"ls", "."}

	}

	if argv[1] == "--help" || argv[1] == "-h" {

		helpLs := new(HelpLs)

		fmt.Println(helpLs.help())

	} else {

		info, err := ioutil.ReadDir(argv[1][:])

		if err != nil {

			fmt.Println("no this dictionary")
		}

		for _, f := range info {

			length := f.Size()

			if ok := f.IsDir(); ok {

				color.Blue("%s", f.Name())

			} else {

				if length > 1024*1024 {

					fmt.Println(f.Name(), "\t", fmt.Sprintf("%.2f", float64(length)/(1024*1024)), "Mb")

				} else {

					fmt.Println(f.Name(), "\t", fmt.Sprintf("%.1f", float64(length)/1024), "Kb")

				}
			}
		}
	}
}
