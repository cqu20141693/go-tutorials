package main

import (
	"fmt"
	"math"
)

/*Constants
constant 关键字用于声明常量变量，必须进行初始化
有全局常量，局部常量
*/
const global string = "globalConstant"

type Color int

// 第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1,一下可以简写
// 枚举
const (
	RED    Color = iota // 0
	ORANGE              // 1
	YELLOW              // 2
	GREEN               // ..
	BLUE
	INDIGO
	VIOLET // 6
)

func Constants() {

	fmt.Println("............Constants............")
	fmt.Println(global)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

	fmt.Println(RED, ORANGE, YELLOW, GREEN, BLUE, INDIGO, VIOLET)
}

/*
Variables
var关键字变量声明,支持单个变量声明，变量初始化，多变量声明
：= 语法 声明并初始化变量
_ 符号表示匿名变量，不需要使用的数据
*/
func Variables() {
	fmt.Println("............Variables............")
	var e int
	fmt.Println(e)

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var (
		name = "jarcper"
		age  = 26
	)
	fmt.Println(name, age)

	var d = true
	fmt.Println(d)

	f := "apple"
	fmt.Println(f)

	// 匿名变量
	_, _ = GetInfo()

}

func GetInfo() (string, int) {
	return "cc", 26
}

/*
Values
go 数据类型：
数字：byte,int8，unit8...int64,unit64,float32,float64,complex
bool
字符串：
数组：
集合：
*/
func Values() {
	fmt.Println("............Values............")
	var boolVar bool
	fmt.Println("bool 默认值：", boolVar)
	var byteVar byte
	fmt.Println("byte 默认值：", byteVar)
	var int8Var int8
	fmt.Println("int8 默认值：", int8Var)
	var int16Var int16
	fmt.Println("int16 默认值：", int16Var)
	var uint8Var uint8
	fmt.Println("uint8Var 默认值：", uint8Var)
	var str string
	fmt.Println("str 默认值：", str)
}
