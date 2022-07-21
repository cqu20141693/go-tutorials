package main

import (
	"fmt"
	"strings"
)

/*init
初始化函数在main函数运行前执行
*/
func init() {
	fmt.Println("如果一个包导入了其他包，则首先初始化导入的包。\n然后初始化当前包的常量。\n接下来初始化当前包的变量。\n最后，调用当前包的 init() 函数。")
}

/*Function
func function_name( [parameter list] ) [return_types] {
   函数体
}

支持无参函数，多参函数，可变参函数，参数传递值拷贝（map,chan,切片默认引用），可以通过* 引用传递
支持无返回值，单个返回值，多返回值


defer 函数： 在函数return前执行
panic 函数：
recover 函数：
闭包 函数：
递归函数：

builtin
len 函数
cap 函数
delete 函数
append函数：
copy 函数：

*/
func Function() {
	fmt.Println("........ Function .........")
	noParameterAndReturn()
	fmt.Printf("input=%s result=%s ", "cqu", singleParameterAndReturn("cqu"))
}
func noParameterAndReturn() {
	fmt.Println("define no parameter and no return")
}
func singleParameterAndReturn(input string) (upper string) {
	fmt.Println("define single parameter and single return")
	return strings.ToUpper(input)
}
