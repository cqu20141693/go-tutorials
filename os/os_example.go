package main

import (
	"fmt"
	"github.com/cqu20141693/go-tutorials/io/tool"
	"os"
	"strings"
)

/*
	os 包提供了系统层面的功能函数
*/

func main() {
	args := os.Args
	fmt.Println("os Args 变量", args)
	who := "Alice "
	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Good Morning", who)
	//  os.Open(path)
	//os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0666)
	//OpenFile 函数有三个参数：文件名、一个或多个标志（使用逻辑运算符 “|” 连接），使用的文件权限
	if !tool.ReadFileDataLineByLine("D:\\go-project\\go-demo\\src\\lib\\os\\os_example.go") {
		return
	}

	testStartProcess()

	//程序处于错误状态时
	os.Exit(1)

	// os.Stdin 标准输入，屏幕是标准输出 os.Stdout；os.Stderr 用于显示错误信息，大多数情况下等同于 os.Stdout
	// 结合bufio使用
	//stdin := os.Stdin

	//os.O_RDONLY：只读
	//os.O_WRONLY：只写
	//os.O_CREATE：创建：如果指定文件不存在，就创建该文件。
	//os.O_TRUNC：截断：如果指定文件已存在，就将该文件的长度截为 0

}

/*
os 包有一个 StartProcess 函数可以调用或启动外部系统命令和二进制可执行文件；它的第一个参数是要运行的进程，第二个参数用来传递选项或参数，第三个参数是含有系统环境基本信息的结构体。
这个函数返回被启动进程的 id（pid），或者启动失败返回错误。
exec 包中也有同样功能的更简单的结构体和函数；主要是 exec.Command(name string, arg ...string) 和 Run()

*/
func testStartProcess() {
	/* Linux: */
	env := os.Environ()
	procAttr := &os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}
	// 1st example: list files
	pid, err := os.StartProcess("/bin/ls", []string{"ls", "-l"}, procAttr)
	if err != nil {
		fmt.Printf("Error %v starting process!", err) //
		os.Exit(1)
	}
	fmt.Printf("The process id is %v", pid)

}
