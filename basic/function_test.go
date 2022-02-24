package basic

import (
	"fmt"
	"testing"
)

/*
	Go 里面有三种类型的函数：
	普通的带有名字的函数
	匿名函数或者 lambda 函数
	方法 ：需要定义接收者

	除了 main ()、init () 函数外，其它所有类型的函数都可以有参数与返回值
	函数被调用的基本格式如下
	pack1.Function(arg1, arg2, …, argn)

	Go 里面函数重载是不被允许的
	函数可以作为类型定义变量，也可以作为函数参数

	任何一个有返回值（单个或多个）的函数都必须以 return 或 panic

	函数参数：
	按值传递：Go 默认使用按值传递来传递参数，也就是传递参数的副本
	引用传递： 传递给函数的是一个指针，通过对指针指向地址的值的修改修改数据
	在函数调用时，像切片（slice）、字典（map）、接口（interface）、通道（channel）这样的引用类型都是默认使用引用传递

	命名返回值：尽量使用命名返回值：会使代码更清晰、更简短，同时更加容易读懂
	命名返回值作为结果形参（result parameters）被初始化为相应类型的零值，当需要返回的时候，我们只需要一条简单的不带参数的 return 语句。
	需要注意的是，即使只有一个命名返回值，也需要使用 () 括起来

	空白符：空白符用来匹配一些不需要的值，然后丢弃掉


	变长参数：函数的最后一个参数是采用 ...type 的形式,长度可以为 0
	同一个类型的变长参数
	不同类型的变长参数（参数总数已知）：可以使用结构体
	不同类型不定长参数：使用空接口 values ...interface{}


	// defer 和追踪 ，参考defer_test.go
	关键字 defer 允许我们推迟到函数返回之前（或任意位置执行 return 语句之后）一刻才执行某个语句或函数
	类似于面向对象编程语言 Java 中finally 语句块
	当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）
	关闭文件流：defer file.Close()
	解锁一个加锁的资源：mu.Lock()  defer mu.Unlock()
	打印最终报告： printFooter()
	关闭数据库链接： defer disconnectFromDB()

	函数作为参数
	函数的返回值和另外一个函数的入参一致
	函数可以作为其它函数的参数进行传递，然后在其它函数内调用执行，一般称之为回调

	函数作为返回值

	闭包：匿名函数
	匿名函数函数不能够独立存在
	匿名函数可以保存函数的地址到变量中
	一个闭包(匿名函数)继承了函数所声明时的作用域。这种状态（作用域内的变量）都被共享到闭包的环境中，因此这些变量可以在闭包中被操作，直到被销毁
	匿名函数可以进行直接调用 func() int { ... } (),花括号 {} 涵盖着函数体，最后的一对括号表示对该匿名函数的调用。
	关键字 defer 经常配合匿名函数使用，它可以用于改变函数的命名返回值
	匿名函数还可以配合 go 关键字来作为 goroutine 使用
*/

func TestAnonymousFunction(t *testing.T) {

	//

	sum := 0
	// 匿名函数直接调用
	func() {
		for i := 1; i <= 1e6; i++ {
			sum += i
		}
	}()
	invokeAnonymousFunc()

	defer func() {
		fmt.Println("can modify return value")
	}()

	go func() {
		fmt.Println("start a goroutine")
	}()

	// 测试闭包（匿名函数）使用定义域内的变量，变量会跟随函数的生命周期
	var f = Adder()
	fmt.Println(f(1), " - ")
	fmt.Println(f(20), " - ")
	fmt.Println(f(300))
}

func Adder() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}
func invokeAnonymousFunc() {
	for i := 0; i < 4; i++ {
		g := func(i int) { fmt.Printf("%d ", i) } //此例子中只是为了演示匿名函数可分配不同的内存地址，在现实开发中，不应该把该部分信息放置到循环中。
		g(i)
		fmt.Printf(" - g is of type %T and has value\n", g)
	}
}
func TestFuncParameter(t *testing.T) {

	noParamFunc()

	sum := addTwoNums(1, 2)
	fmt.Println("run addTwoNums sum=", sum)

	slice := []int{1, 2, 3, 4}
	sum = SumSlice(slice)
	fmt.Println("run SumSlice slice =", sum)

	// 没有命令返回值
	x2, x3 := runNotCmdReturn(1)
	fmt.Println("runNotCmdReturn x2,x3=", x2, x3)

	// 引用传递
	input := 12
	referenceParam(&input)
	fmt.Println("reference param input=", input)

	// 变长参数
	sum = SumVariableLengthParam(1, 2, 3, 4)
	fmt.Println("SumVariableLengthParam=", sum)
	sum = SumVariableLengthParam(1, 2, 3, 4, 5)

	// struct 参数,默认也是拷贝赋值
	employee := Employee{"gowb", 25, 10000.0}
	CheckEmployee(employee)
	fmt.Println("struct :CheckEmployee ", employee)
	CheckEmployee2(&employee)
	fmt.Println("&struct :CheckEmployee ", employee)

	//slice,map,interface{},channel 参数
	ints := []int{1, 2, 3}
	m := map[string]int{"age": 23}
	checkReference(ints, m)
	fmt.Println("checkReference ", ints, m)
}

// 函数作为参数
func TestFuncParam(t *testing.T) {
	fmt.Println("test function parameter")

	callback(1, Add)
}

func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

func callback(y int, f func(int, int)) {
	f(y, 2) // this becomes Add(1, 2)
}

func checkReference(ints []int, m map[string]int) {
	for i := range ints {
		ints[i]++
	}
	for k := range m {
		m[k] += 3
	}
}

func CheckEmployee2(employee *Employee) {
	if employee.age > 24 {
		fmt.Println("employee name=", employee.name)
		employee.salary = employee.salary + 1000
		return
	}
}

func CheckEmployee(employee Employee) {
	if employee.age > 24 {
		fmt.Println("employee name=", employee.name)
		employee.salary = employee.salary + 1000
		return
	}
}

func SumVariableLengthParam(values ...int) (sum int) {
	for _, value := range values {
		sum += value
	}
	return
}

func referenceParam(input *int) {
	*input = (*input) * 2
}

func runNotCmdReturn(input int) (int, int) {
	x2 := 2 * input
	x3 := 3 * input
	return x2, x3
}

func SumSlice(slice []int) (sum int) {
	// test 空白符
	for _, value := range slice {
		sum += value
	}
	for i := range slice {
		sum += slice[i]
	}
	// 返回命名返回值 sum
	return
}

func addTwoNums(i int, i2 int) (sum int) {
	return i + i2
}

func noParamFunc() {
	fmt.Println("run noParamFunc")
}
