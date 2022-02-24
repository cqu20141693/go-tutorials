package basic

import (
	"fmt"
	"testing"
)

/*
	return 返回值其实是对返回变量的返回
	如果再defer 中对返回值变量操作会改动返回值
*/
func TestReturn(t *testing.T) {
	fmt.Println("return-------", test())
	fmt.Println("return value-------", testReturnValue())
	fmt.Println("return defer value-------", testReturnDeferValue())
	fmt.Println()
}

func testReturnDeferValue() (n string) {
	defer func() {
		n = "gowb"
	}()
	return
}

func testReturnValue() (n string) {
	s := "hello wujt"
	fmt.Println(n)
	return s
}

/*
	返回变量为n 返回类型是string
	直接return 则默认是n的值
*/
func test() (n string) {
	n = "hello"
	return
}
