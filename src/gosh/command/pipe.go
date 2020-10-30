package command

//2020.10.26增加对内置命令 `>` 以及`>>`的支持 系统命令已经支持该操作,具体实现过程尚不知晓
//增加对`|`的支持
//当写到这里时发现前面代码设计的不够好，或许在添加一些新的东西时，很多代码需要重构

//对与有输出内容的命令。不能单纯的输出到os.Stdout中
//充分运用os.Stdin os.Stderr
//应当设计一些接口 否则很局限 学学unix的设计


type unixLike interface {

	// ">"覆盖 && ">>"追加
	toFile(name string)

	// "|" 输出重定向
	pipe()

}


