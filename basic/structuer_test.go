package basic

import (
	"fmt"
	"runtime"
	"strings"
	"testing"
)

/*
 	if-else 结构
	关键字 if 和 else 之后的左大括号 { 必须和关键字在同一行，如果你使用了 else-if 结构，
	则前段代码块的右大括号 } 必须和 else-if 关键字在同一行
	当 if 结构内有 break、continue、goto 或者 return 语句时，Go 代码的常见写法是省略 else 部分

	switch 结构
变量选择模式
switch var1 {
    case val1:
        ...
    case val2:
        ...
    default:
        ...
}
	变量 var1 可以是任何类型，而 val1 和 val2 则可以是同类型的任意值。类型不被局限于常量或整数，但必须是相同的类型；
	或者最终结果为相同类型的表达式。前花括号 { 必须和 switch 关键字在同一行
	多个可能符合条件的值，使用逗号分割它们，例如：case val1, val2, val3
	一旦成功地匹配到某个分支，在执行完相应代码后就会退出整个 switch 代码块，不需要特别使用 break 语句来表示结束。
	如果在执行完每个分支的代码后，还希望继续执行后续分支的代码，可以使用 fallthrough 关键字来达到目的
	可以使用 return 语句来提前结束代码块的执行
	可选的 default 分支可以出现在任何顺序，但最好将它放在最后


条件模式
switch {
    case condition1:
        ...
    case condition2:
        ...
    default:
        ...
}


条件模式（包含初始化）
switch initialization {
    case condition1:
        ...
    case condition2:
        ...
    default:
        ...
}
*/
var prompt = "Enter a digit, e.g. 3 " + "or %s to quit."

func Init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else { //Unix-like
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
}

func TestIfElse(t *testing.T) {
	bool1 := false
	if bool1 {
		fmt.Printf("The value is true\n")
		return
	} else {
		fmt.Printf("The value is false\n")
	}

	var cond int
	if cond = getCond(); cond < 10 {
		fmt.Println("cond is less than 10")
	}

	fmt.Println(prompt)
	Init()
	fmt.Println(prompt)

}

func getCond() int {
	return 5
}

func TestSwitch(t *testing.T) {
	i := 0
	switch i {
	case 0: // 空分支，只有当 i == 0 时才会进入分支
	case 1:
		invoke(i) // 当 i == 0 时函数不会被调用
	}

	switch i {
	case 0:
		fallthrough
	case 1:
		invoke(i) // 当 i == 0 时函数也会被调用
	}

	switch i {
	case 0, -1:
		invoke(i)
		fallthrough
	case 1:
		invoke(i)
		fallthrough
	default:
		fmt.Println("default")
	}

	ConditionMode(-1)
	ConditionMode(0)
	ConditionMode(2)

	ConditionAndInitMode(1)
	ConditionAndInitMode(6)
	ConditionAndInitMode(13)

}

func ConditionAndInitMode(month int) {
	switch season := Season(month); {
	case strings.EqualFold("Winter", season) || strings.EqualFold("Spring", season):
		fmt.Println("开心的日子")
		// 不需要break
	case strings.EqualFold("Summer", season) || strings.EqualFold("Autumn", season):
		fmt.Println("收获的日子")
	default:
		fmt.Println("放纵的日子")
	}
}

func ConditionMode(num1 int) {
	switch {
	case num1 < 0:
		fmt.Println("Number is negative")
	case num1 > 0 && num1 < 10:
		fmt.Println("Number is between 0 and 10")
	default:
		fmt.Println("Number is 10 or greater")
	}
}

func Season(month int) string {
	switch month {
	case 12, 1, 2:
		return "Winter"
	case 3, 4, 5:
		return "Spring"
	case 6, 7, 8:
		return "Summer"
	case 9, 10, 11:
		return "Autumn"
	}
	return "Season unknown"
}

func invoke(value int) {
	fmt.Println("invoked ", value)
}
