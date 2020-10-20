package command

import (

	"fmt"
	"github.com/fatih/color"
	"io/ioutil"

)

func Ls(argv []string) {

	if len(argv) == 1 {

		argv = []string{"ls", "."}

	}

	info, err := ioutil.ReadDir(argv[1][:])

	if err != nil {

		fmt.Println("no this dictionary")
	}

	for _, f := range info {

		length := f.Size()

		if ok := f.IsDir(); ok {

			color.Blue("%s", f.Name())

		} else {

			if length > 1024*1024 {

				fmt.Println(f.Name(), " ", fmt.Sprintf("%.2f", float64(length)/(1024*1024)), "Mb")

			} else {

				fmt.Println(f.Name(), " ", fmt.Sprintf("%.1f", float64(length)/1024), "Kb")
			}
		}

	}
}
