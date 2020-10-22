package gosh

import (

    "fmt"
    "github.com/carmark/pseudo-terminal-go/terminal"
    "github.com/nj-jay/go-shell/gosh/command"
    "github.com/spf13/viper"
    "log"
    "os"
    "os/user"
    "strings"

)

type ShellPrompt struct {


}

type User struct {

    Username string

    Homedir string

    CurrentPath string

    Hostname string
}



var (

    reader string

    shellPrompt *ShellPrompt

    Usr User

    term *terminal.Terminal

)

func Initialize() {

    term, _ = terminal.NewWithStdInOut()

    defer term.ReleaseFromStdInOut()

    reader, _=term.ReadLine()

}

func (*User) GetUserInfo() {

    currentUser, _ := user.Current()

    Usr.Username = currentUser.Username

    Usr.Hostname, _ = os.Hostname()

    Usr.CurrentPath, _ = os.Getwd()

    Usr.Homedir = currentUser.HomeDir

}


func (*ShellPrompt) PrintHeader() {

    themeName := shellPrompt.CreateConfig()

    LoadTheme(themeName)

}

//将用户的输入进行格式化 如　" ls   ..  " -> ["ls", ".."]
func (*ShellPrompt) FormatInput() ([]string, error) {

    input := strings.Replace(reader, "\n\n", "", -1)

    input = strings.TrimRight(input, "\r\n")

    formatInputFields := strings.Fields(input)

    return formatInputFields, nil
}

func (*ShellPrompt) CreateConfig() string {

    Usr.GetUserInfo()

    configFile := Usr.Homedir + "/config.toml"

    if !command.PathExists(configFile) {

        viper.SetConfigName("config")

        viper.SetConfigType("toml")

        viper.AddConfigPath(Usr.Homedir)

        viper.Set("app_name", "go-shell")

        viper.Set("Theme.theme", "yes")

        err := viper.SafeWriteConfig()

        if err != nil {

            log.Fatal("write config failed: ", err)

        }
    }

    viper.SetConfigName("config")

    viper.SetConfigType("toml")

    viper.AddConfigPath(Usr.Homedir)

    err := viper.ReadInConfig()

    if err != nil {

        fmt.Println("error")
    }

    return viper.GetString("Theme.theme")
}





