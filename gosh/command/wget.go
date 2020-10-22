package command

import (

	"fmt"
	"gopkg.in/cheggaaa/pb.v1"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

)

const (

	ContextWget = `
	wget url [url]
	
	wget url			根据url下载文件到当前目录下
	wget url url			多线程下载
`

)
type HelpWget struct {


}

func (*HelpWget) help() string {

	return ContextWget
}

var wg sync.WaitGroup

//支持多线程下载
func Wget(argv []string) {

//打印时间

	if argv[1] == "--help" || argv[1] == "-h" {

		helpWget := new(HelpWget)

		fmt.Println(helpWget.help())

	} else {

		for _, url := range argv[1:] {

			wg.Add(1)

			go func(url string) {

				defer wg.Done()
				getMultiple(url)

			}(url)

		}

		wg.Wait()

	}



}

func getMultiple(url string) {

	res, _ := http.Get(url)

	fmt.Print("\n已发出" + res.Proto + "请求, 正在等待回应...")

	//判断Status
	if res.StatusCode == 404 {

	fmt.Println(res.Status)

	os.Exit(0)

	} else if res.StatusCode == 200 {

		fmt.Println(res.Status)

		index := strings.LastIndex(url, "/")

		//当前目录创建文件
		downFile, _ := os.Create(url[index+1:])

		defer downFile.Close()

		length := res.Header.Get("Content-Length")

		lengthKb, _ := strconv.ParseFloat(length, 64)

		fmt.Println("长度:	", "[", length, "byte", "]", ">--", fmt.Sprintf("%.2f", lengthKb/1024), "Kb", "	[", res.Header.Get("Content-Type"), "]")

		fmt.Println("正在保存至:", url[index+1:])

		//创建一个进度条
		bar := pb.New(int(lengthKb)).SetUnits(pb.U_BYTES_DEC).SetRefreshRate(time.Millisecond * 10)

		// 显示下载速度
		bar.ShowSpeed = true

		// 显示剩余时间
		bar.ShowTimeLeft = true

		// 显示完成时间
		bar.ShowFinalTime = true

		bar.SetWidth(80)

		bar.Start()

		writer := io.MultiWriter(downFile, bar)

		//写到f中
		io.Copy(writer, res.Body)

		//确保finish
		bar.Finish()

		//关闭连接
		_ = res.Close

		log.Println("已保存", url[index+1:])

		}
}