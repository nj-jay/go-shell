package command

import (
	"fmt"
	"os/exec"
)

/*
暂时使用notepad记事本代替
 */

const (

	ContextGedit = `
	gedit hello.c		用文本编辑器打开文本文件hello.c
	gedit d:/test/hello.c	打开d:/test目录下的文本文件hello.c	
`

)
type HelpGedit struct {


}

func (*HelpGedit) help() string {

	return ContextGedit
}


func Gedit(argv []string) {

	if len(argv) == 1 {

		cmd := exec.Command("cmd.exe", "/c", "notepad.exe")

		err := cmd.Run()

		if err != nil {

			fmt.Println("gedit: command not found")

		}

		fmt.Println("warning! gedit filename")

	} else if len(argv) == 2 && (argv[1] == "--help" || argv[1] == "-h") {

		helpGedit := new(HelpGedit)

		fmt.Println(helpGedit.help())

	} else {

		cmd := exec.Command("cmd.exe", "/c", "notepad.exe", argv[1])

		err := cmd.Run()

		if err != nil {

			fmt.Println("gedit: command not found")

		}

	}

}