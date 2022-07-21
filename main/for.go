package main

import (
	"fmt"
	"reflect"
)

/*For
for 关键字用于循环，可以没有表达式，无限循环，
只有一个条件表达式，while替代
三个表达式，初始化，执行条件，后续行为（initial/condition/after）
支持break关键字退出循环

特别地: go 中条件不需要(), 包括for,if else,switch,select
go 没有while
*/
func For() {
	fmt.Println("............For............")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}

/*Range
for range 可以遍历数组、切片、字符串、map 及通道（channel），for range 语法上类似于其它语言中的 foreach 语句
*/
func Range() {
	fmt.Println("..........Range............")
	// range string
	str := "hello world."
	for index, char := range str {
		// rune int32
		fmt.Println("index", index, "value", char, "type", reflect.TypeOf(char))
	}
	for index := range str {
		// unit8 byte 字符
		char := str[index]
		fmt.Println("index", index, "value", char, "type", reflect.TypeOf(char))
	}
	// range array
	arr := [...]byte{'1', '2', '3'}
	for index, bytes := range arr {
		fmt.Println("index", index, "value", bytes, "type", reflect.TypeOf(bytes))
	}
	for index := range arr {
		char := arr[index]
		fmt.Println("index", index, "value", char, "type", reflect.TypeOf(char))
	}

	// range slice
	sli := []rune{'中', '国'}
	for index, int32Var := range sli {
		fmt.Println("index", index, "value", int32Var, "type", reflect.TypeOf(int32Var))
	}
	for index := range sli {
		runeVar := arr[index]
		fmt.Println("index", index, "value", runeVar, "type", reflect.TypeOf(runeVar))
	}

	// range map
	strMap := map[int]bool{1: true, 2: true, 3: true}
	for key, value := range strMap {
		fmt.Println("key", key, "key-type", reflect.TypeOf(key), "value", value, "type", reflect.TypeOf(value))
	}

	for key := range strMap {
		value := strMap[key]
		fmt.Println("key", key, "key-type", reflect.TypeOf(key), "value", value, "type", reflect.TypeOf(value))
	}

	// range chan
	c := make(chan int)

	go func() {

		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()
	// 当chan close时推出遍历
	for v := range c {
		fmt.Println(v)
	}
	intChan := make(chan int)
	go func() {

		intChan <- 1
		intChan <- 2
		intChan <- 3
		close(intChan)
	}()
	for {
		var isBreak bool
		select {
		case v, ok := <-intChan:
			if ok {
				fmt.Println(v)
			} else {
				fmt.Println("chan is closed")
				isBreak = true
			}
		}
		if isBreak {
			break
		}
	}

	// range
	fmt.Println("..........Range............ End")
}
