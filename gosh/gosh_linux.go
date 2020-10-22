package gosh

import (

    "fmt"
    "github.com/nj-jay/go-shell/gosh/command"
    "io"
    "os"
)

//for循环不断获得用户输入的命令

func Run() {


    //direct to $HOME
    command.Cd([]string{"cd"})

    //清屏
    command.Cls([]string{"cls"})

    for {


        shellPrompt.PrintHeader()

        Initialize()

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