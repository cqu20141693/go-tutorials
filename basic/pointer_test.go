package basic

import (
	"fmt"
	"testing"
)

/*
	Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址
	var p *type 定义指针变量
	符号 * 可以放在一个指针前，如 *p，那么它将得到这个指针指向地址上所存储的值
*/
func TestPointer(t *testing.T) {
	var i1 = 5
	fmt.Printf("An integer: %d, its location in memory: %p\n", i1, &i1)
	var intP *int
	intP = &i1
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)
}
