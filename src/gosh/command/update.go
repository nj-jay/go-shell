package command

import (
	"fmt"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

/*
更新go-shell命令
再服务器创建一个资源库
从服务器端进行下载
https://gosh.nj-jay.com
 */

func Update(argv []string) {

	if len(argv) == 1 {

		//Wget([]string{"wget", "-d", "D:/go-shell/bin", "https://picture.nj-jay.com/pwd.exe"})
		//fmt.Println(GetAllRemoteFile())
		path := "D:/go-shell/bin/"

		RemoteFileInfo := GetAllRemoteFile()

		//保存下载好的程序的文件名以及时间
		//LocalAllFileInfo := make(map[string]string)

		viper.SetConfigName("update")
		viper.SetConfigType("toml")
		viper.AddConfigPath("D:/go-shell")

		for index, value := range RemoteFileInfo {

			if !PathExists(path + index) {

				//fmt.Println(index)

				Wget([]string{"wget", "-d", "D:/go-shell/bin", "https://gosh.nj-jay.com/" + index})

				viper.Set(index, value)
				//viper.SafeWriteConfig()
				//
				//if err != nil {
				//
				//	fmt.Println("write config failed: ", err)
				//}
				err := viper.WriteConfig()

				if err != nil {

					fmt.Println("write config failed: ", err)
				}
				fmt.Println("go-shell命令已经是最新")

			} else if PathExists(path + index) {

				//可能得用本地文件进行存储时间吧 update.toml
				err := viper.ReadInConfig()

				if err != nil {

					fmt.Println("error")
				}

				//test
				//fmt.Println(viper.GetString("go-shell.exe"))
				if value != viper.GetString(index) {

					//如果更新了已经存在的程序 就更新
					fmt.Println("检测到更新 正在更新程序!")
					//[putty]
					//  exe = "16042167455097931"
					Wget([]string{"wget", "-d", "D:/go-shell/bin", "https://gosh.nj-jay.com/" + index})

					RemoteFileInfo = GetAllRemoteFile()

					for index, value := range RemoteFileInfo {

						viper.Set(index, value)

						err := viper.WriteConfig()

						if err != nil {

							fmt.Println("write config failed: ", err)
						}
					}



				}
			}
		}
	}

	fmt.Println("go-shell命令已经是最新")
}

func GetAllRemoteFile() map[string]string {

	//map必须初始化才能使用
	RemoteAllFileInfo := make(map[string]string)

	accessKey := "0BXxOytiWWOySUiEbg-t8_08c0tvPpwTwDA6Ivfn"

	secretKey := "ewA76OHGBEz43vQg-gCqOHd_5paEurmSEzFdx-dz"

	mac := qbox.NewMac(accessKey, secretKey)

	cfg := storage.Config{

		UseHTTPS: true,

	}

	bucketManager := storage.NewBucketManager(mac, &cfg)

	bucket := "go-shell"

	limit := 1000

	prefix := ""

	delimiter := ""

	//初始列举marker为空
	marker := ""

	for {

		entries, _, nextMarker, hasNext, err := bucketManager.ListFiles(bucket, prefix, delimiter, marker, limit)

		if err != nil {

			fmt.Println("list error,", err)

			break

		}

		//print entries
		for _, entry := range entries {

			//fmt.Println(entry.Key, entry.PutTime)
			//保存文件的名字与时间戳
			RemoteAllFileInfo[entry.Key] = cast.ToString(entry.PutTime)

		}

		if hasNext {

			marker = nextMarker

		} else {

			break

		}
	}

	return RemoteAllFileInfo
}






