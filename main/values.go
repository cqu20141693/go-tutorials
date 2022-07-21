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

/*Types
在使用int和 uint类型时，不能假定它是32位或64位的整型，而是考虑int和uint可能在不同平台上的差异。
Go语言中的字符有两种：用''
uint8类型，或者叫byte型，代表ASCII 码的一个符号；
rune类型，代表一个UTF-8 符号；
*/
func Types() {
	fmt.Println("............Types............")
	// 断言一个接口对象（i）里不是 nil，并且接口对象（i）存储的值的类型是 T，如果断言成功，就会返回值给 t，如果断言失败，就会触发 panic。
	var data interface{} = "hello"
	str := data.(string)
	fmt.Println(str)
	// 不会panic,断言失败，此时t 为 T 的零值。
	var i interface{} // nil
	v, ok := i.(interface{})
	fmt.Println(v, ok)

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int8:
			fmt.Println("I'm a int8")
		case int16:
			fmt.Println("I'm a int16")
		case int32:
			fmt.Println("I'm a int32/rune")
		case int64:
			fmt.Println("I'm a int64")
		case int:
			fmt.Println("I'm an int")
		case uint8: //type byte=unit8
			fmt.Println("I'm a uint8/byte")
		case uint16:
			fmt.Println("I'm a uint16")
		case uint32:
			fmt.Println("I'm a uint32")
		case uint64:
			fmt.Println("I'm a uint64")
		case uint:
			fmt.Println("I'm an uint")
		case float32:
			fmt.Println("I'm a float32")
		case float64:
			fmt.Println("I'm an float64")
		case string:
			fmt.Println("I'm an string")
		case complex64:
			fmt.Println("I'm an complex64")
		case complex128:
			fmt.Println("I'm an complex128")
		case uintptr:
			fmt.Println("I'm an uintptr,无符号整型，用于存放指针")

		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
	whatAmI(nil)

}

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
	fmt.Println("type byte=uint8 默认值：", byteVar)
	var int8Var int8
	fmt.Println("int8 默认值：", int8Var)
	var int16Var int16
	fmt.Println("int16 默认值：", int16Var)
	var uint8Var uint8
	fmt.Println("uint8Var 默认值：", uint8Var)
	var str string
	fmt.Println("str 默认值：", str)
}
