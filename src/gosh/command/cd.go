package command

import (

	"fmt"
	"os"
	"os/user"

)

const (

	 ContextCd = `
	cd [path]
			cd			进入家目录
			cd ..			进入上一级目录
			cd work			进入当前目录的work目录
			cd d:			进入d盘
`

)
type HelpCd struct {


}

func (*HelpCd) help() string {

	return ContextCd

}

// Cd changes the shell's directory
func Cd(argv []string) {

	if len(argv) == 1 {

		//if formatInput == "cd"
		//return ~/$HOME
		currentUser, _ := user.Current()

		err := os.Chdir(currentUser.HomeDir)

		if err != nil {

			fmt.Println("go-shell: cd: 没有这个目录")

		}

	} else if argv[1] == "--help" || argv[1] == "-h" {

		helpCd := new(HelpCd)

		fmt.Println(helpCd.help())

	} else {

		err := os.Chdir(argv[1])

		if err != nil {

			fmt.Println("cd: 没有那个文件或目录")
		}
	}
}