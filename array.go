// 当前项目的包名
package main

// 导入其他的包
import (
	"fmt"
)

// 由 main 函数作为程序的入口点启动
func main() {
	// 初始化一个长度为 2 的数组
	a := [2]int{1, 1}

	// 初始化一个长度为 10 的数组，并将索引为 1 的值赋值为 2
	b := [10]int{1: 2}

	// 当不知道数组长度为多少时，可以使用 ...
	c := [...]int{1, 2, 3, 4, 5, 6}

	// 当不知道数组长度为多少时，可以使用 ...
	d := [...]int{0: 1, 2: 3}

	fmt.Println(a);
	fmt.Println(b);
	fmt.Println(c);
	fmt.Println(d);

	var arr1 [5]int

	for i := 0; i < len(arr1); i++ {
		arr1[i] = i * 2
	}

	for i := 0; i < len(arr1); i++ {
		fmt.Printf("Array at index %d is %d\n",i,arr1[i])
	}

	// IDIOM
	arr2 := [...]string{"a", "b", "c", "d"}

	for i := range arr2 {
		fmt.Println("Arrary item", i, "is", arr2[i])
	}

	//多维数组
	const(
		WIDTH = 5
		HEIGHT = 3
	)

	var screen [WIDTH][HEIGHT]int

	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			screen[x][y] = x
		}
		fmt.Printf("Array at index %d\n",screen)
	}
	
}
