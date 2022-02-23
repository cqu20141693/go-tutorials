package main // package 是 Go 语言声明包的关键字，main 是要声明的包名。在 Go 语言中 main 包是一个特殊的包，代表你的 Go 语言项目是一个可运行的应用程序，而不是一个被其他项目引用的库。

import "fmt" // import 是 Go 语言的关键字，表示导入包的意思，这里我导入的是 fmt 包

/*
package main 和 main 函数这两个核心部分， package main 代表的是一个可运行的应用程序，而 main 函数则是这个应用程序的主入口。
*/

var a string
var done bool

func setup() {
	a = "hello, world"
	done = true
}
func main() { //其中 func 是 Go 语言的关键字，表示要定义一个函数或者方法的意思，main 是函数名，() 空括号表示这个 main 函数不接受任何参数。在 Go 语言中 main 函数是一个特殊的函数，它代表整个程序的入口
	fmt.Println("hello,world！") //Println 函数是属于包 fmt 的函数

	go setup()
	for !done {
	}
	print(a)
}
