package command

import (
	"fmt"
	"os"
)

const (

	ContextTouch = `
	touch filenamePath(文件的路径)

	touch filename				在当前目录下创建文件
	touch ../filename			在上一级目录下创建文件
	touch d:/test.txt			在d盘下创建文件test.txt
`

)
type HelpTouch struct {


}

func (*HelpTouch) help() string {

	return ContextTouch
}

func Touch(argv []string) {

	if len(argv) == 1 {

		fmt.Println("touch:缺少了文件操作数")

	} else if argv[1] == "--help" || argv[1] == "-h" {

		helpTouch := new(HelpTouch)

		fmt.Println(helpTouch.help())

	} else {

		if !PathExists(argv[1]) {

			f, err := os.Create(argv[1])

			//记得close掉
			//否则go-shell未关闭的时候无法删除该文件

			defer f.Close()

			if err != nil {

				fmt.Println("touch: 创建文件失败")

			}

		} else if PathExists(argv[1]) {

			f, err := os.Open(argv[1])

			defer f.Close()

			if err != nil {

				fmt.Println("go-shell: touch: 没有这个文件")

			}

			//Todo
			//修改文件的最后修改时间 使用touch命令解决



		}
	}
}