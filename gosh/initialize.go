package gosh

import (

	"bufio"
	"fmt"
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

}



var (

	reader *bufio.Reader

	shellPrompt *ShellPrompt

	Usr User

	HostName string

)

func initialize() {

	//when working, direct to HOME directory
	command.Cd([]string{"cd"})

	//get user's input
	reader = bufio.NewReader(os.Stdin)

}

func (*User) GetUserInfo() {

	currentUser, _ := user.Current()

	//这个名字更合理
	Usr.Username = currentUser.Name

	Usr.CurrentPath, _ = os.Getwd()

	Usr.Homedir = currentUser.HomeDir

}


func (*ShellPrompt) printHeader() {

	themeName := shellPrompt.createConfig()

	LoadTheme(themeName)

}

//将用户的输入进行格式化 如　" ls   ..  " -> ["ls", ".."]
func (*ShellPrompt) formatInput() ([]string, error) {

	input, err := reader.ReadString('\n')

	input = strings.TrimRight(input, "\n")

	formatInputFields := strings.Fields(input)

	return formatInputFields, err
}

func (*ShellPrompt) createConfig() string {

	Usr.GetUserInfo()

	configFile := Usr.Homedir + "/.config.toml"

	historyFile := Usr.Homedir + "/.go-shell.history"

	//创建.go-shell.history
	if !command.PathExists(historyFile) {

		f, _ := os.Create(historyFile)

		defer f.Close()

	}

	if !command.PathExists(configFile) {

		viper.SetConfigName(".config")

		viper.SetConfigType("toml")

		viper.AddConfigPath(Usr.Homedir)

		viper.Set("app_name", "go-shell")

		viper.Set("Theme.theme", "jay")

		err := viper.SafeWriteConfig()

		if err != nil {

			log.Fatal("write config failed: ", err)

		}
	}

	viper.SetConfigName(".config")

	viper.SetConfigType("toml")

	viper.AddConfigPath(Usr.Homedir)

	err := viper.ReadInConfig()

	if err != nil {

		fmt.Println("error")
	}

	HostName = viper.GetString("hostname")

	if HostName == "" {

		HostName = "go-shell"

	}
	return viper.GetString("Theme.theme")

}