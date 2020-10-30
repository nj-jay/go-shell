package command

import (

	"fmt"
	"os"

)


const (

	ContextExit = `
	exit			退出go-shell
`

)
type HelpExit struct {


}

func (*HelpExit) help() string {

	return ContextExit
}

func Exit(argv []string) {

	if len(argv) == 1 {

		os.Exit(0)

	} else if argv[1] == "--help" || argv[1] == "-h" {

		helpExit := new(HelpExit)

		fmt.Println(helpExit.help())

	}

}