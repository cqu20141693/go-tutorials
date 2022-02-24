package basic

import (
	"fmt"
	"testing"
)

/*
	label 标记作用的范围
	goto
	一般不推荐使用标签和goto，程序不方便理解
	goto使用过程中，推荐定义顺序的标签，跳过多少程序；一个比较好的场景是死循环中出现问题，跳出循环，但是也可以使用break跳出，除非多层循环

*/
func TestLabel(t *testing.T) {

	// 作用范围是外层循环
LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}

	for i := 0; i <= 5; i++ {
		// 作用为当前层for循环
	LABEL2:
		for j := 0; j <= 5; j++ {
			if j == 3 {
				// 当前作用可以用break替代
				continue LABEL2
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}

	// 作用范围是外层循环
LABEL3:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				break LABEL3
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}
	// 作用范围是外层循环

	for i := 0; i <= 5; i++ {
	LABEL4:
		for j := 0; j <= 5; j++ {
			if j == 3 {
				// 等价不需要label
				break LABEL4
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}

}

func TestGoto(t *testing.T) {
	i := 0
HERE:
	print(i)
	i++
	if i == 5 {
		return
	}
	goto HERE
}
