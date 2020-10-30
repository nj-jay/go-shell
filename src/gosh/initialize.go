package gosh

import (

	"bufio"
	"fmt"
	"github.com/nj-jay/go-shell/gosh/command"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strings"

)


func init() {

	//修改窗口标题
	cmd := exec.Command("cmd.exe", "/c", "title=go-shell v1.0.0")

	cmd.Stdin = os.Stdin

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr

	cmd.Run()

}

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

	configPath := "D:/go-shell"

	configFile := "D:/go-shell/config.toml"

	historyFile := "D:/go-shell/go-shell.history"

	envPath := "D:/go-shell/bin"

	command.Mkdir([]string{"mkdir", configPath})

	command.Mkdir([]string{"mkdir", envPath})

	//创建.go-shell.history
	if !command.PathExists(historyFile) {

		f, _ := os.Create(historyFile)

		defer f.Close()

	}

	if !command.PathExists(configFile) {

		viper.SetConfigName("config")

		viper.SetConfigType("toml")

		viper.AddConfigPath(configPath)

		viper.Set("app_name", "go-shell")

		viper.Set("Theme.theme", "jay")

		err := viper.SafeWriteConfig()

		if err != nil {

			log.Fatal("write config failed: ", err)

		}
	}

	viper.SetConfigName("config")

	viper.SetConfigType("toml")

	viper.AddConfigPath(configPath)

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