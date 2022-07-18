package main

import (
	"fmt"
	"time"
)

/*SwitchCase

switch case 关键子用于枚举条件处理
switch 支持无变量，一个变量
case 支持常量值条件，表达式条件， case 代码段执行完
fallthrough 继续下一个case
default 表示无匹配case时执行逻辑
*/
func SwitchCase() {
	fmt.Println("............SwitchCase............")
	i := 1
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one and fallthrough")
		fallthrough
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon use default")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

}

/*IfElse
if else 关键字用于条件分支语句
支持单个表达式判断执行分支
支持两个表达式，初始化表达式，判断表达式实现条件判断
*/
func IfElse() {
	fmt.Println("............IfElse............")
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

}
