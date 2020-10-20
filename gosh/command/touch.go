package command

import (

	"fmt"
	"os"

)

func Touch(argv []string) {

	if len(argv) == 1 {

		fmt.Println("touch:缺少了文件操作数")

	} else {

		if !PathExists(argv[1]) {

			_, err := os.Create(argv[1])

			if err != nil {

				fmt.Println("touch: 创建文件失败")

			}

		}

	}

}