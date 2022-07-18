package practice

import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutineFor(t *testing.T) {
	list := [...]int{1, 3, 5, 7, 9, 10}
	// 各个goroutine中输出的i,v值都是for range循环结束后的i, v最终值
	for i, v := range list {
		go func() {
			fmt.Println(i, v)
		}()
	}
	time.Sleep(3 * time.Second)
	print("利用函数参数传递")
	for i, v := range list {
		go func(index, value int) {
			fmt.Println(index, value)
		}(i, v)
	}
	time.Sleep(3 * time.Second)
}
