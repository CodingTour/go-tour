package main

import (
	"fmt"
	"os"
)

/*
这是一个多行注释。
并且注释需要紧挨着程序实体，即变量、常量、函数、结构体以及接口。
*/
func main() {
	// 这是一个单行注释。
	// 与Java的main方法不同，Go的main函数无法直接获取输入参数。
	var args []string = os.Args
	fmt.Printf("args: %v\n", args)

	fmt.Println("Hello World!")
}