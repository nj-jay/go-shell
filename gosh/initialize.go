package gosh

import (

    "bufio"
    "fmt"
    "os"
    "os/user"
)

const (

    shellLoginText = `
go-shell by nj-jay v0.0.1
    `

)

var (

    reader *bufio.Reader

    shellPrompt *ShellPrompt

)

type ShellPrompt struct {


}

func initialize() {

    fmt.Println(shellLoginText)

    reader = bufio.NewReader(os.Stdin)

}


func (*ShellPrompt) PrintHeader(){

    CurrentUser, _ := user.Current()

    username := CurrentUser.Username

    hostname, _ := os.Hostname()

    CurrentPath, _ := os.Getwd()

    fmt.Print("# " + username + "@" + hostname +
        " in " + CurrentPath + "\n" + "$ ")
}


