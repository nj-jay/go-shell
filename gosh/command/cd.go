package command

import (
	"fmt"
	"os"
)

// Cd changes the shell's directory
func Cd(argv []string) error {

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

	} else {

		/*
		if formatInput == "cd work"
		first get CurrentPath
		then change path
		*/

		wd, err := os.Getwd()

		if err != nil {

			return err
		}

		err = os.Chdir(wd + "/" + argv[1])

		if err != nil {

			fmt.Println("cd: 没有那个文件或目录")
		}

	}

	return nil
}