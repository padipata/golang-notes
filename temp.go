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
				break LABEL3
			}
		}

}


