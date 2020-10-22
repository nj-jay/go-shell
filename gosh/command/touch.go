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
	touch /path/to/filename			在/path/to目录下创建文件
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

			_, err := os.Create(argv[1])

			if err != nil {

				fmt.Println("touch: 创建文件失败")

			}

		}

	}
}