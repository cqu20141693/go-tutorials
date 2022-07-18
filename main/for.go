package main

import "fmt"

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
