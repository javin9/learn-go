## Go 语言学习笔记

最近在学习 go 语言，走了很多弯路，本仓库记录了日常测试的代码片段

### go 安装

```bash
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

### goland 安装

1.创建新项目的时候，选 go module，第一个已经不在用了
2.keymap 设置快捷键
3.Tools -> File wathcer -> + 保存自动排版【已有】 4.一个文件件存在多个 main 函数，上面文本框，选择-Edit Congifratin -> Run kind -> File

### 初始化环境

https://beego.gocn.vip/beego/zh/developing/bee/env.html

```bash
#go语言安装主根目录
export GOROOT=/usr/local/go #替换你的目录
#GOPATH 是自己的go项目路径，自定义设置
export GOPATH=/Users/ding/go_workspace #替换你的目录
#GOBIN 当我们使用go install命令编译后并且安装的二进制程序目录
export GOBIN=$GOPATH/bin
# 启用 Go Modules 功能
export GO111MODULE=on
# 配置 GOPROXY 环境变量。你可以换成你喜欢的 proxy
export GOPROXY=https://goproxy.cn,direct
# 加入环境变量中
export PATH=$PATH:$GOROOT/bin:$GOBIN
```

#go 语言安装主根目录

export GOROOT=/usr/local/go

替换你的目录#GOPATH 是自己的 go 项目路径，自定义设置 export GOPATH=/Users/ding/go_workspace #替换你的目录#GOBIN 当我们使用 go install 命令编译后并且安装的二进制程序目录 export GOBIN=$GOPATH/bin# 启用 Go Modules 功能export GO111MODULE=on# 配置 GOPROXY 环境变量。你可以换成你喜欢的 proxyexport GOPROXY=https://goproxy.cn,direct# 加入环境变量中export PATH=$PATH:$GOROOT/bin:$GOBIN

### 创建文件夹

mkdir -p $GOPAH/src/github.com/rupid/learn-gin

### 进入 learn-gin 文件夹

cd $\_

### 初始化 mod

go mod init

### 拉去 gin 的模块

go get -v github.com/gin-gonic/gin@v1.7

### validator

v8 版本是`binding`，v9 以后就改成了`validate`
https://pkg.go.dev/github.com/go-playground/validator@v9.31.0+incompatible#section-readme

### 练习

- 类型转换是强制的
- varname:=134 这种格式只能放到 func 内，不能全局用，全局只能用 var varname int
- //const 不需要转换 看代码 数值可作为各种类型使用
- iota 的使用 \_ 跳过
- if 后面的赋值，括号外访问不到
- switch 会自己 break；fallthrough

### 参考资料

- [Go Module 构建模式如何对依赖包进行升级和降级](http://www.go-day.cn/detail/6.html)
- [govendoe](https://blog.csdn.net/weixin_33389183/article/details/112099594)
- [go-gin-api](https://github.com/xinliangnote/go-gin-api)
- [gin_scaffold](https://github.com/e421083458/gin_scaffold)
