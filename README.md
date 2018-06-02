# Golang-notes
Golang takes notes from the beginning to the actual combat

目录
===

<!-- TOC -->

- [环境安装](#环境安装)
- [Hello world](#Hello world)

<!-- /TOC -->
---


## 环境安装
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

