package command

import (

	"fmt"
	"os"

)

const (

	 ContextCd = `
	cd [path]
			cd			进入家目录
			cd ..			进入上一级目录
			cd work			进入当前目录的上一级目录
			cd /home/jay		进入/home/jay目录
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
		_ = os.Chdir(os.Getenv("HOME"))

	} else if argv[1][0:1] == "/" {

		/*
		if formatInput == "cd /etc"
		 */

		err := os.Chdir(argv[1])

		if err != nil {

			fmt.Println("cd: 没有那个文件或目录")
		}

	} else if argv[1][0:1] == "~" {

		// if formatInput == "cd ~/work"
		err := os.Chdir(os.Getenv("HOME") + argv[1][1:])

		if err != nil {

			fmt.Println("cd: 没有这个文件或目录")
		}
	} else if argv[1] == "--help" || argv[1] == "-h" {

		helpCd := new(HelpCd)

		fmt.Println(helpCd.help())

	} else {

		/*
		if formatInput == "cd work"
		first get CurrentPath
		then change path
		*/

		wd, _ := os.Getwd()

		err := os.Chdir(wd + "/" + argv[1])

		if err != nil {

			fmt.Println("cd: 没有那个文件或目录")
		}

	}

}