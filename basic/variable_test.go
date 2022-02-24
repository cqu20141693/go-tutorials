package basic

import (
	"fmt"
	"testing"
)

/*
	当一个变量被声明之后，系统自动赋予它该类型的零值：
	int 为 0，float 为 0.0，bool 为 false，string 为空字符串，指针为 nil。
	记住，所有的内存在 Go 中都是经过初始化的。

*/

var a, b *int

var a1 int
var b1 bool
var str1 string

var (
	a2   int
	b2   bool
	str2 string
)

func TestVariable(t *testing.T) {
	fmt.Println(a, b)
	fmt.Println(a1, b1, str1)
	fmt.Println(a2, b2, str2)
}
