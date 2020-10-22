package command

import (

	"fmt"
	"os"
	"os/exec"
	"runtime"

)

var (

	System string = runtime.GOOS

)

const (

	ContextCls = `
	cls			清屏
`

)
type HelpCls struct {


}

func (*HelpCls) help() string {

	return ContextCls
}

func Cls(argv []string) {

	if len(argv) == 1 {

		switch System {

		case "linux":

			cmd := exec.Command("bash", "-c", "clear")

			//必备不能少
			cmd.Stdin = os.Stdin

			cmd.Stdout = os.Stdout

			cmd.Stderr = os.Stderr

			err := cmd.Run()

			if err != nil {

				fmt.Println("cls: command not found")

			}

		case "windows":

			cmd := exec.Command("cls")

			cmd.Stdin = os.Stdin

			cmd.Stdout = os.Stdout

			cmd.Stderr = os.Stderr

			err := cmd.Run()

			if err != nil {

				fmt.Println("cls: command not found")

			}

		}

	} else if argv[1] == "--help" || argv[1] == "-h" {

		helpCls := new(HelpCls)

		fmt.Println(helpCls.help())

	}


}
