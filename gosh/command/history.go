package command

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
)

/*

记录100条命令
如果超过100条
删除第一个
保证数量是100
并将数据保存在.go-shell.history文件中

 */

var (

	currentUser, _ = user.Current()

	historyFile = currentUser.HomeDir + "/.go-shell.history"

)

//先读文件,打印文件中的内容
func History(argv []string) {

	if len(argv) == 1 {

		//	line := LineHistory()
		//
		//	fmt.Println("line", line)
		//
		//	if line > 50 {
		//
		//		//if line ==51
		//
		//		buf, _ := ioutil.ReadFile(historyFile)
		//
		//		index := strings.Index(string(buf), "\n")
		//
		//		fmt.Println(string(buf[index:]))
		//
		//	} else {
		//
		//			Cat([]string{"cat", "-n", historyFile})
		//	}
		//}
	}

	Cat([]string{"cat", "-n", historyFile})

}

func WriteHistory(argv []string) {

	//要给写的权限
	//有个坑, os.Open()这个方法没有写的权限!!!!!!!!!!!!
	//并且是追加的形式，否则每次将被清除!!!!!!!!!!!
	f, err := os.OpenFile(historyFile, os.O_WRONLY | os.O_APPEND, 0666)

	defer f.Close()

	if err != nil {

		fmt.Println("文件打开失败")

	}

	outputWriter := bufio.NewWriter(f)

	var readerInput string

	for _, buf := range argv {

		readerInput += buf + " "

	}

	_, _ = outputWriter.WriteString(readerInput + "\n")

	_ = outputWriter.Flush()

}

//bug
//func LineHistory() int {
//
//	//判断文件的行数 如果超过100删除第一行
//
//	var line int = 1
//
//	f, err := os.OpenFile(historyFile, os.O_RDONLY, 0666)
//
//	defer f.Close()
//
//	if err != nil {
//
//		fmt.Println("go-shell: history: 没有.go-shell.history这个文件")
//
//		return 0
//
//	}
//
//	reader := bufio.NewReader(f)
//
//	for  {
//
//		_, err := reader.ReadString('\n')
//
//		if err == io.EOF {
//
//			fmt.Println()
//
//			return line
//
//
//		}
//
//		line = line + 1 //行数加1
//
//	}
//
//}

