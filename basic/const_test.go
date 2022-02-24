package basic

import "testing"

/*
	常量使用关键字 const 定义，用于存储不会改变的数据。
	存储在常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。
	常量也允许使用并行赋值的形式：
	常量还可以用作枚举,
	iota 可以被用作枚举值,每遇到一次 const 关键字，iota 就重置为 0
*/
const age int = 25

const beef, two, food = "eat", 2, "veg"
const (
	Monday, Tuesday, Wednesday = 1, 2, 3
	Thursday, Friday, Saturday = 4, 5, 6
)
const (
	Unknown = 0
	Female  = 1
	Male    = 2
)

type Color int

// 第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1,一下可以简写
const (
	RED    Color = iota // 0
	ORANGE              // 1
	YELLOW              // 2
	GREEN               // ..
	BLUE
	INDIGO
	VIOLET // 6
)

func TestConst(t *testing.T) {

}
