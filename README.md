# Golang-notes
Golang takes notes from the beginning to the actual combat

目录
===

<!-- TOC -->

- [基础入门](#基础入门)
    - [环境安装](#环境安装)
    - [Hello world](#hello-world)
    - [基础变量](基础变量)
    - [基础运算](基础运算)

<!-- /TOC -->


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

// 导入包
import (
	io "fmt"
)

// 常量赋值
// const PI = 3.14

// 没有赋值变量时，结果为零值
var PI int
// var PI int = 123

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

### 基础运算
```go
// 当前项目的包名
// 当前项目的包名
package main

// 导入其他的包
import (
	"fmt"
)

const (
	a int = 1
	b string = "abc"
)

// 由 main 函数作为程序的入口点启动
func main() {
	c:= len(b);
	for i := 0; i < c; i++ {
		fmt.Println(i)
	}

	a := 1
	switch a {
	case 0:
		fmt.Println("a=0")
	case 1:
		fmt.Println("a=1")
	case 2:
		fmt.Println("a=2")
	default:
		fmt.Println("None")
	}

	// 跳出到 LABEL1 之外，可以是任意循环
	LABEL1:
		for {
			for i := 0; i < 10; i++ {
				fmt.Println(i)
				if i > 3 {
					break LABEL1
				}
			}
		}
			
	// 跳出到 LABEL2 之外，LABEL2 需要写在后面
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			if i > 3 {
				goto LABEL2
			}
		}
	LABEL2:

	// 跳出到 LABEL3 之外，必须是有限循环
	LABEL3:
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			if i > 3 {
				continue LABEL3
			}
		}

}
```


