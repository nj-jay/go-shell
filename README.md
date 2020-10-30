# go-shell

## 简介

采用go语言编写的shell

仿照linux/GNU系统中bash

重写linux中的各种内置命令并加以一些个人认为好用的命令

如linux中的ls, cd, mkdir, touch等

使用go语言进行重写并内置到go-shell中

it works well on windows!

It will be better in future!

## go-shell配置文件

本项目初衷完全采用linux bash理念

如bash的配置文件为$HOME/.bashrc

go-shell的配置文件为d:/go-shell/config.toml

修改配置文件d:/go-shell/config.toml可以go-shell的一些参数

如主题默认的主题是yes, 可以更改为jay

目前提供两种主题 yes 和 jay

<img src="https://picture.nj-jay.com/2020.10.22.21.38.png" style="zoom: 80%;" />

<img src="https://picture.nj-jay.com/2020.10.22.21.39.png" style="zoom: 67%;" />

## 设计思路

本项目设计的初衷是:

>  window下的cmd不是很好用于是设计一个与linux相同的命令行解释器, 能够与系统进行交互。操作各种文件而不需要使用各种编程语言去操作。另外由于go语言能够直接编译成各平台可直接运行的二进制程序，所以该项目可以很方便的进行使用.



在初步完善该项目后发现:A

> 完成一个项目需要设计很多东西，当代码比较混乱时需要考虑使用｀设计模式｀，需要设计｀各种错误类型｀，能够让该程序稳定的运行而不至于程序崩溃。在设计各种功能的时候考虑使用不同的数据结构，能够让代码简洁优雅。如map类型, struct, interface接口等等。在逐步完善的过程中，能够将所学到的一些知识完美的结合起来真是一件无比开心的事情！！！
>
> 另外该项目完整做下来是一个巨大的工程，因为要重写linux中的各种命令，虽然可以调用linux上的命令，但是我想做的是跨平台使用，所以想使用go语言重写各种命令。这不是一个痛苦的事情，相反能够让我更加熟悉go语言.

##  使用go-shell

运行程序后

输入`help`查看当前版本支持的命令

输入cd --help或者cd -h可查看cd命令的具体使用

输入update命令更新go-shell命令

> 更新的来源：gosh.nj-jay.com
>
> 我将所有编译好的二进制文件发布在服务器上
>
> 用户只需要输入update即可更新命令

另外为了用户更好程序使用go-shell

go-shell兼容cmd

## 支持用户自定义命令

为了增强go-shell的可使用性

您可以写一些您需要的程序供自己使用

只需要把编译好的二进制文件放在d:/go-shell/bin目录下即可

## Example

```shell
./go-shell

# jay@nj in /home/jay
$ ls

# 结果
Projects
SoftWare
baidunetdiskdownload
fast
me
snap
test
work

# jay@nj in /home/jay
$ cd work
# jay@nj in /home/jay/work
$ ls

#结果
bin
pkg
src

# jay@nj in /home/jay/work
$ mkdir test

#结果
bin
pkg
src
test
```

## 后期目标

后期会逐步实现linux各种命令并内置到go-shell中

同时也会设计不同的主题让go-shell变得更加优雅

同时我会设计一些好用的命令让go-shell变得更加强大

本项目也将由我个人长期维护更新

欢迎提issue

有兴趣的朋友也可联系我一起进行开发
