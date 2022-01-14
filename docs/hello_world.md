# 你好，世界！

1. [Downloads](https://go.dev/dl/)，下载一个适合的版本，并配置环境变量。
2. 创建一个工程目录，如go-tour，并切换到该目录
3. 初始化工程
```bash
go mod init github.com/CodingTour/go-tour
```
简而言之，mod指的是Modules，即Go的包管理器，类似Java的Maven。

更多细节可以参考[Tutorial: Create a Go module](https://go.dev/doc/tutorial/create-module)

为了日后方便使用，国内的用户还需要设置一下代理，具体可以参考[Goproxy.cn](https://goproxy.cn/)

4. Coding

```go
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
```

5. 运行

Go可以使用go run命令运行：
```bash
go run hello_world.go
```

也可以先编译：

```bash
go build hello_world.go
```

然后直接运行：

```bash
./hello_world
```


[源码](../examples/hello_world/hello_world.go)

在这里还需要说明几个Go语言的特点：

- 包名(main)可与目录名（hello_world）不一致
- 作为程序的入口，包名必须为“main” 