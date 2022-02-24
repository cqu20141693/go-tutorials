package main

import (
	"fmt"
)

/*
	//  从控制台读取输入
	Scanln 扫描来自标准输入的文本，将空格分隔的值依次存放到后续的参数内，直到碰到换行。
	Scanf 与Scanln类似，除了 Scanf 的第一个参数用作格式字符串，用来决定如何读取.
	Sscan 和以 Sscan 开头的函数则是从字符串读取，除此之外，与 Scanf 相同
*/
var (
	firstName, lastName, s string
	i1                     int
	f1                     float32
	input                  = "56.12 / 5212 / Go"
	format                 = "%f / %d / %s"
)

func main() {
	TestScan()
}

func TestScan() {
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName, &lastName)
	fmt.Printf("Hi %s %s!\n", firstName, lastName) // Hi Chris Naegels
	fmt.Scanf("%s %s", &firstName, &lastName)
	fmt.Printf("Hi %s %s!\n", firstName, lastName) // Hi Chris Naegels
	fmt.Sscanf(input, format, &f1, &i1, &s)
	fmt.Println("From the string we read: ", f1, i1, s)
	// 输出结果: From the string we read: 56.12 5212 Go

	// 数据写入文件
	// fmt.Fprintf(outputFile, “Some test data.\n”)
}
