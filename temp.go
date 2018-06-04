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


	LABEL1:
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			if i > 3 {
				break LABEL1
			}
		}

}


