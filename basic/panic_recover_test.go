package basic

import (
	"fmt"
	"testing"
	"time"
)

/*
使用 panic 和 recover 关键字时遇到的现象，部分现象也与上一节分析的 defer 关键字有关：

panic 只会触发当前 Goroutine 的 defer；
recover 只有在 defer 中调用才会生效；
panic 允许在 defer 中嵌套多次调用；

*/

// 测试panic 在当前goroutine 内有效
func Test1(t *testing.T) {
	defer println("in test")
	go func() {
		defer println("in goroutine")
		panic("")
	}()
	time.Sleep(1 * time.Second)
}

// 失效的崩溃恢复
func Test2(t *testing.T) {
	defer fmt.Println("in test")
	if err := recover(); err != nil {
		fmt.Println(err)
	}

	panic("unknown err")
}

// 嵌套崩溃
func Test3(t *testing.T) {
	defer func() {
		fmt.Println("in test")
		if err := recover(); err != nil {
			fmt.Println("in test recover")
		}
	}()
	defer func() {
		defer func() {
			panic("panic again and again")
		}()
		panic("panic again")
	}()

	panic("panic once")
}
