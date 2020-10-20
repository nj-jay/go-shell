package command

import (

	"fmt"
	"os"

)

func Mkdir(argv []string) {

	//判断目录存不存在
	b := PathExists(argv[1])

	//if not exist
	if !b {

		err := os.Mkdir(argv[1], 0777)

		if err != nil {

			fmt.Println("mkdir: 无法创建目录: 权限不够")

		}
	}

}

func PathExists(path string) bool {

	_, err := os.Stat(path)

	if err == nil {

		return true
	}
	if os.IsNotExist(err) {

		return false
	}

	return false
}