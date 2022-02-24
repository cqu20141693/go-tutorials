package basic

import (
	"fmt"
	"testing"
)

/*
	complex，real,imag: 创建和操作复数
	cap: 返回array元素个数，slice 元素个数，channel buffer量，nil 返回0
	close: 关闭close,对于值为nil的channel或者对同一个channel重复close, 都会panic, 关闭只读channel会报编译错误。
	len: 用于返回某个类型的长度或数量（字符串、数组、切片、map 和管道）
	new：new(type)分配内存，用于值类型和用户定义的类型，如自定义结构,new (T) 为每个新的类型 T 分配一片内存，初始化为 0 并且返回类型为 * T 的内存地址：
	返回一个指向类型为 T，值为零值的地址的指针，它适用于值类型如数组和结构体；它相当于 &T{}
	make：make(type)用于内置引用类型（切片、map 和管道）
	append : 切片添加
	copy: 复制切片
	panic，recover：两者均用于错误处理机制
	print,println: 底层打印函数，在部署环境中建议使用 fmt 包
*/

func TestBuildInFunction(t *testing.T) {
	var strings []string
	strings = append(strings, "gowb")
	fmt.Println("append ", strings)

	capacity := cap(strings)
	fmt.Println(capacity)
	fmt.Println()
	testClose()
	fmt.Println()
	var x complex128 = complex(1, 2)       // 1+2i
	var y complex128 = complex(3, 4)       // 3+4i
	fmt.Println("complex", x*y)            // "(-5+10i)"
	fmt.Println("complex real", real(x*y)) // "-5"
	fmt.Println("complex imag", imag(x*y)) // "10"
}

func testClose() {
	ch1 := make(chan int)
	ch2 := make(chan bool)
	// close channel
	go write(ch1)
	go read(ch1, ch2)

	<-ch2
}

func read(ch1 chan int, ch2 chan bool) {
	for {
		v, ok := <-ch1
		if ok {
			fmt.Printf("read a int is %d\n", v)
		} else {
			ch2 <- true
		}
	}

}

func write(ch chan int) {

	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
