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

泛型接口约束只能用于参数： 方法参数，泛型方法，

*/

/*
方法中直接使用约束，不利用接口声明的约束
*/
func add[T int | float64](a, b T) T {
	return a + b
}

// 别名支持
func stringaAdd[T ~string](x, y T) T {
	return x + y
}

type MyString string

// 这里的 any 并非泛型的约束，而是类型
func testAny(x any) any {
	return x
}

/**
测试泛型方法，参数
*/
func TestConstraintParameter(t *testing.T) {

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
	fmt.Println(testAny("a"))
}

/**
方法泛型
*/
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

type myInt int

func (m myInt) DoSomething() {
	fmt.Println("do something")
}

type myInterface interface {
	DoSomething()
}

type MyStruct[T myInterface] struct {
	Value T
}

func (receiver MyStruct[T]) printValue() {
	fmt.Println(receiver.Value)
}

// 一个泛型接口(关于泛型接口在后半部分会详细讲解）
type IPrintData[T int | float32 | string] interface {
	Print(data T)
}

// 一个泛型通道，可用类型实参 int 或 string 实例化
type MyChan[T int | string] chan T

/**
测试泛型类型(Generic type)
申明泛型类型时使用类型形参， 使用泛型类型变量时使用类型实参
匿名结构体不支持泛型
~ 支持底层相同数据类型的泛型形参:使用 ~ 时有一定的限制：
~后面的类型不能为接口
~后面的类型必须为基本类型
*/
func TestConstraintsField(t *testing.T) {

	// 自定义容器泛型类型（结构体，接口）
	selfDefineGeneric()

	// map 类型泛型
	mapGeneric()

	// slice类型泛型
	sliceGeneric()

	// 泛型receiver
	receiverGeneric()
	// 泛型类型断言
	genericTypeJudge()

	genericAlias()
}

/**
Go新增了一个符号 ~ ，在类型约束中使用类似 ~int 这种写法的话，就代表着不光是 int ，所有以 int 为底层类型的类型也都可用于实例化。
*/
type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Uint interface {
	~uint | ~uint8 | ~uint16 | ~uint32
}
type Float interface {
	~float32 | ~float64
}

func genericAlias() {

	type Slice[T Int | Uint | Float] []T

	var s Slice[int] // 正确
	fmt.Println("genericAlias Slice[int]：", s)
	type MyInt int
	var s2 Slice[MyInt] // MyInt底层类型是int，所以可以用于实例化
	fmt.Println("genericAlias Slice[MyInt]：", s2)
	type MyMyInt MyInt
	var s3 Slice[MyMyInt] // 正确。MyMyInt 虽然基于 MyInt ，但底层类型也是int，所以也能用于实例化
	fmt.Println("genericAlias Slice[MyMyInt]：", s3)
	type MyFloat32 float32 // 正确
	var s4 Slice[MyFloat32]
	fmt.Println("genericAlias Slice[MyFloat32]：", s4)
}

// 这里类型约束使用了空接口，代表的意思是所有类型都可以用来实例化泛型类型 Queue[T] (关于接口在后半部分会详细介绍）
// type any = interface{}
type Queue[T any] struct {
	elements []T
}

// 将数据放入队列尾部
func (q *Queue[T]) Put(value T) {
	q.elements = append(q.elements, value)
}

// 从队列头部取出并从头部删除对应数据
func (q *Queue[T]) Pop() (T, bool) {
	var value T
	if len(q.elements) == 0 {
		return value, true
	}

	value = q.elements[0]
	q.elements = q.elements[1:]
	return value, len(q.elements) == 0
}

// 队列大小
func (q *Queue[T]) Size() int {
	return len(q.elements)
}

/**
泛型类型不支持类型断言
*/
func genericTypeJudge() {
	var q1 Queue[int] // 可存放int类型数据的队列
	q1.Put(1)
	q1.Put(2)
	q1.Put(3)
	fmt.Println("put queue[int]:", q1)
	q1.Pop() // 1
	q1.Pop() // 2
	q1.Pop() // 3
	fmt.Println("pop queue[int]:", q1)

	var q2 Queue[string] // 可存放string类型数据的队列
	q2.Put("A")
	q2.Put("B")
	q2.Put("C")

	pop, b := q2.Pop() // "A"
	if !b {
		fmt.Println("generic not support type assert .(string)/.(type) ", pop)
		//  not support 类型断言
		//s := pop.(string)
		//switch pop.(type) {

		//}
	}
	q2.Pop() // "B"
	q2.Pop() // "C"

	//var q3 Queue[struct{Name string}]
	//var q4 Queue[[]int] // 可存放[]int切片的队列
	//var q5 Queue[chan int] // 可存放int通道的队列
	//var q6 Queue[io.Reader] // 可存放接口的队列

}

type MySlice[T int | float32] []T

/**
通用的方法处理泛型类型
*/
func (s MySlice[T]) Sum() T {
	var sum T
	for _, value := range s {
		sum += value
	}
	return sum
}

func receiverGeneric() {

	myslice := make(MySlice[int], 10)
	fmt.Println("receiver generic:", myslice)
}

func selfDefineGeneric() {
	vars := make([]MyStruct[myInt], 5)
	vars[0] = MyStruct[myInt]{Value: myInt(1)}
	for _, m := range vars {
		m.printValue()
	}
}

func mapGeneric() {
	// MyMap类型定义了两个类型形参 KEY 和 VALUE。分别为两个形参指定了不同的类型约束
	// 这个泛型类型的名字叫： MyMap[KEY, VALUE]
	type MyMap[KEY int | string, VALUE float32 | float64] map[KEY]VALUE

	// 用类型实参 string 和 flaot64 替换了类型形参 KEY 、 VALUE，泛型类型被实例化为具体的类型：MyMap[string, float64]
	var a MyMap[string, float64] = map[string]float64{
		"jack_score": 9.6,
		"bob_score":  8.4,
	}
	fmt.Println("MyMap[string, float64]:", a)
}

type Slice[T int | float32 | float64] []T

func sliceGeneric() {
	// 这里传入了类型实参int，泛型类型Slice[T]被实例化为具体的类型 Slice[int]
	var a Slice[int] = []int{1, 2, 3}
	fmt.Printf("Type Name: %T", a) //输出：Type Name: Slice[int]

	// 传入类型实参float32, 将泛型类型Slice[T]实例化为具体的类型 Slice[string]
	var b Slice[float32] = []float32{1.0, 2.0, 3.0}
	fmt.Printf("Type Name: %T", b) //输出：Type Name: Slice[float32]

}
