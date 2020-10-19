package gosh

import (
    "bufio"
    "fmt"
    "github.com/nj-jay/go-shell/gosh/command"
    "os"
    "os/user"
    "strings"
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

    command.Cd([]string{"cd"})


}


func (*ShellPrompt) PrintHeader() {

    CurrentUser, _ := user.Current()

    username := CurrentUser.Username

    hostname, _ := os.Hostname()

    CurrentPath, _ := os.Getwd()

    fmt.Print("# " + username + "@" + hostname +
        " in " + CurrentPath + "\n" + "$ ")
}

func (*ShellPrompt) formatInput() ([]string, error) {

    input, err := reader.ReadString('\n')

    input = strings.TrimRight(input, "\r\n")

    formatInputFields := strings.Fields(input)

    return formatInputFields, err
}
