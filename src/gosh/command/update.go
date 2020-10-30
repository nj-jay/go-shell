package command

/*
更新go-shell命令
再服务器创建一个资源库
从服务器端进行下载
https://gosh.nj-jay.com
 */

func Update(argv []string) {

	if len(argv) == 1 {

		Wget([]string{"wget", "-d", "D:/go-shell/bin", "https://picture.nj-jay.com/pwd.exe"})
	}

}





