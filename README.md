# Golang-notes
Golang takes notes from the beginning to the actual combat

目录
===

<!-- TOC -->

- [基础入门](#基础入门)
    - [环境安装](#环境安装)
    - [Hello world](#hello-world)
    - [基础变量](基础变量)

<!-- /TOC -->
---

## 基础入门

### 环境安装
```shell
#安装包下载地址（mac使用pkg格式）
https://golang.google.cn/dl/

#检查安装是否成功
go env

#将 /usr/local/go/bin 目录添加至PATH环境变量：
export PATH=$PATH:/usr/local/go/bin

注意：MAC 系统下你可以使用 .pkg 结尾的安装包直接双击来完成安装，安装目录在 /usr/local/go/ 下。

#安装成功
go
```

### Hello world
> touch hello.go

```go
package main

import (
	"fmt"
)

func main(){
	fmt.Println("hello world")
}
```

> 运行及打包

```go
go run hello.js
go build hello.js
```

### 基础变量

```go
// 当前项目的包名
package main

// 导入其他的包
import (
	io "fmt"
)

// 常量赋值
// const PI = 3.14

// 没有赋值变量时，结果为零值
var PI int

// 全局变量
var name = "gopher"

// 一般变量
type newType int

// 结构类型的声明
type  gopher struct{

}

// 接口的声明
type golang interface{

}

// 由 main 函数作为程序的入口点启动
func main() {
	io.Println(PI)
}
```

***注意：golang中，使用首字母大小写来决定该 变量 常量 类型 接口 结构 函数 是否被外部调用
（小写：public，大写：private）***


