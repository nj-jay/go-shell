package main

import (

	"fmt"
	"os"

)

//pwd.exe

func main() {

	//获得用户当前的path

	path, _ := os.Getwd()

	fmt.Println(path)


}