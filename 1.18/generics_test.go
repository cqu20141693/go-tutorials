package __18

import (
	"fmt"
	"testing"
)

/*
// 没有任何约束
func add[T any](x, y T) T
// 约束 Addble (需要单独定义)
func add[T Addble](x, y T) T
// 约束允许 int 或 float64 类型
func add[T int|float64](x, y T) T
// 约束允许底层类型是 string 的类型（包括 string 类型）
func add[T ~string](x, y T) T


*/

/*
方法中直接使用约束，不利用接口
*/
func add[T int | float64](a, b T) T {
	return a + b
}
func stringaAdd[T ~string](x, y T) T {
	return x + y
}

type MyString string

// 这里的 any 并非泛型的约束，而是类型
func test(x any) any {
	return x
}

func TestConstraint(t *testing.T) {

	// 利用| 构建约束，可读性差
	fmt.Println(add(1, 2))
	fmt.Println(add(1.2, 2.3))

	// 基于某类型定义新类型，有时可能希望泛型约束是某类型的所有衍生类型
	// 使用type定义新类型，利用~ 表示基类型的衍生类型
	var x string = "ab"
	var y MyString = "cd"
	fmt.Println(stringaAdd(x, x))
	fmt.Println(stringaAdd(y, y))

	// any  interface{}的衍生类型
	fmt.Println(test("a"))
}

func SumIntsOrFloats[K comparable, V int | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// 申明泛型约束
type Number interface {
	int | float64
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
func TestConstraints(t *testing.T) {
	ints := map[string]int{"1": 1, "2": 2}
	floats := map[string]float64{"1": 1.0, "2": 2.0}
	fmt.Printf("泛型计算结果，Ints 结果%v: Floats 结果: %v\n", SumIntsOrFloats[string, int](ints), SumIntsOrFloats[string, float64](floats))
	fmt.Printf("泛型计算结果，Ints 结果%v: Floats 结果: %v\n", SumNumbers[string, int](ints), SumNumbers[string, float64](floats))

}
