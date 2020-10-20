package gosh

import (

	"fmt"
	"github.com/spf13/cast"
	"time"
)
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

	fmt.Print("# " + Usr.Username + "@" + Usr.Hostname +
		" in " + Usr.CurrentPath + "\n" + "$ ")

}

func (*Theme) Jay() {

	Usr.GetUserInfo()

	Now := time.Now()

	timeNow := "[" + cast.ToString(Now.Hour()) + ":" + cast.ToString(Now.Minute()) +
		":" + cast.ToString(Now.Second()) + "]"

	fmt.Print("# " + Usr.Username + "@" + Usr.Hostname +
		" in " + Usr.CurrentPath + "  " + timeNow +  "\n" + "$ ")

}