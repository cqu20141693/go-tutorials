package basic

import (
	"fmt"
	"testing"
)

/*
	基于计数器的迭代
	for 初始化语句; 条件语句; 修饰语句 {}

	基于条件判断的迭代
	for 条件语句 {}

	无限循环
	i:=0; ; i++ 或 for { } 或 for ;; { }

	可以使用 break 语句（第 5.5 节）或 return 语句直接返回
	在 switch 或 select 语句中（详见第 13 章），break 语句的作用结果是跳过整个代码块，执行后续的代码
	关键字 continue 忽略剩余的循环体而直接进入下一次循环的过程，关键字 continue 只能被用于 for 循环中

	for-range 结构
	可以迭代任何一个集合（包括数组和 map)
	可以获得每次迭代所对应的索引。一般形式为：for ix, val := range coll { }
	val 始终为集合中对应索引的值拷贝，因此它一般只具有只读性质，对它所做的任何修改都不会影响到集合中原有的值
	如果 val 为指针，则会产生指针的拷贝，依旧可以修改集合中的原值
*/

func TestForCounter(t *testing.T) {
	for i := 0; i < 5; i++ {
		fmt.Printf("This is the %d iteration\n", i)
	}
	i := 0
START:
	fmt.Printf("The counter is at %d\n", i)
	i++
	if i < 15 {
		goto START
	}
}

func TestStringLoop(t *testing.T) {
	str := "Go is a beautiful language!"
	fmt.Printf("The length of str is: %d\n", len(str))
	for ix := 0; ix < len(str); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str[ix])
	}
	str2 := "日本語"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for ix := 0; ix < len(str2); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str2[ix])
	}
}

func TestBitWise(t *testing.T) {
	fmt.Println("测试补码")
	for i := 0; i <= 10; i++ {
		fmt.Printf("the complement of %b is: %b\n", i, ^i)
	}
}

func TestForConditionMode(t *testing.T) {

	i := 5
	for i >= 0 {
		i--
		fmt.Printf("the variable is %d \n", i)
	}
}

func TestInfiniteLoop(t *testing.T) {

	var length int = 10
	i := 0
	var str string
	for {
		i++
		if i > length {
			break
		}
		str += "G"
		fmt.Println(str)
	}
}

func TestForRange(t *testing.T) {
	str := "Go is a beautiful language!"
	fmt.Printf("The length of str is: %d\n", len(str))
	for pos, char := range str {
		fmt.Printf("Character on position %d is: %c \n", pos, char)
	}
	fmt.Println()
	str2 := "Chinese: 日本語"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for pos, char := range str2 {
		fmt.Printf("character %c starts at byte position %d\n", char, pos)
	}
	fmt.Println()
	fmt.Println("index int(rune) rune    char bytes")
	for index, r := range str2 {
		fmt.Printf("%-2d      %d      %U '%c' % X\n", index, r, r, r, []byte(string(r)))
	}
}
