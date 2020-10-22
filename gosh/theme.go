package gosh

import (

	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cast"
	"time"

)

//根据配置文件选择主题

type Theme struct {


}

func LoadTheme(theme string) {

	themeName := new(Theme)

	switch theme {

	case "yes":

		themeName.Yes()

	case "jay":

		themeName.Jay()

	}

}

func (*Theme) Yes() {

	Usr.GetUserInfo()

	//fmt.Print("# " + Usr.Username + "@" + Usr.Hostname +
	//	" in " + Usr.CurrentPath + "\n" + "$ ")

	black := color.New(color.FgBlack)

	blueBackground := black.Add(color.BgBlue)

	//redBackground := black.Add(color.BgRed)
	blueBackground.Print(Usr.CurrentPath + " ")

	fmt.Print(" ☞  ")

}

func (*Theme) Jay() {

	Usr.GetUserInfo()

	Now := time.Now()

	timeNow := "[" + cast.ToString(Now.Hour()) + ":" + cast.ToString(Now.Minute()) +
		":" + cast.ToString(Now.Second()) + "]"

	//fmt.Print("# " + Usr.Username + "@" + Usr.Hostname +
	//	" in " + Usr.CurrentPath + "  " + timeNow +  "\n" + "$ ")

	//jay主题
	fmt.Print(color.BlueString("# ") + color.CyanString(Usr.Username) + " @ " + color.GreenString(Usr.Hostname) +
		" in " + color.YellowString(Usr.CurrentPath) + "  " + timeNow + "\n" + color.RedString("$ "))


}