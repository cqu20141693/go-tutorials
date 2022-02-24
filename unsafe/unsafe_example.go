package main

import (
	"fmt"
	"github.com/cqu20141693/go-tutorials/basic"
	"unsafe"
)

/*

 */
func main() {
	employee := new(basic.Employee)
	employee.SetName("gowb")

	sizeof := unsafe.Sizeof(employee)
	fmt.Println("unsafe SizeOf employee ", sizeof)
}
