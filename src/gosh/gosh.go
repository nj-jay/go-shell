package gosh

import (
	"fmt"
	"io"
	"os"
)

func Run() {

	// initialize the go-shell
	initialize()

	////清屏
	//command.Cls([]string{"cls"})

	//设置环境变量

	for {

		shellPrompt.printHeader()

		err := interpret()

		if err == nil {

			continue

		}

		if err == io.EOF {

			_, _ = fmt.Fprintln(os.Stderr, err)

			break
		}

		fmt.Println(err)

	}

}