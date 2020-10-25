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
	_, _ = blueBackground.Print(Usr.CurrentPath + " ")

	fmt.Print(" ☞  ")

}

func (*Theme) Jay() {

	Usr.GetUserInfo()

	Now := time.Now()

	//使用匿名函数
	hour := func() string {

		if len(cast.ToString(Now.Hour())) == 1 {

			return "0" + cast.ToString(Now.Hour())

		}
		return cast.ToString(Now.Hour())
	}()

	minute := func() string {

		if len(cast.ToString(Now.Minute())) == 1 {

			return "0" + cast.ToString(Now.Minute())

		}
		return cast.ToString(Now.Minute())
	}()

	second := func() string {

		if len(cast.ToString(Now.Second())) == 1 {

			return "0" + cast.ToString(Now.Second())

		}
		return cast.ToString(Now.Second())
	}()

	//还是有bug 当时间变为0时，只显示一位
	//timeNow := "[" + cast.ToString(Now.Hour()) + ":" + cast.ToString(Now.Minute()) +
	//	":" + cast.ToString(Now.Second()) + "]"

	timeNow := "[" + hour + ":" + minute +
		":" + second + "]"

	//bug “颜色显示乱码”
	//jay主题
	//fmt.Print(color.BlueString("# ") + color.CyanString(Usr.Username) + " @ " + color.GreenString(Usr.Hostname) +
	//	" in " + color.YellowString(Usr.CurrentPath) + "  " + timeNow + "\n" + color.RedString("$ "))
	//

	red := color.New(color.FgRed)
	cyan := color.New(color.FgCyan)
	green := color.New(color.FgGreen)
	yellow := color.New(color.FgYellow)
	_, _ = red.Print("# ")
	_, _ = cyan.Print(Usr.Username)
	fmt.Print(" @ ")
	_, _ = green.Print(HostName)
	fmt.Print(" in ")
	_, _ = yellow.Print(Usr.CurrentPath)
	fmt.Print("  ")
	fmt.Print(timeNow + "\n")
	_, _ = red.Print("$ ")

}