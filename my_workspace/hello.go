package main

import (
	"fmt"
	"github.com/samber/lo"
	"strconv"
)

/**
https://go.dev/doc/tutorial/workspaces
*/
func main() {
	// uniq : 底层利用map 去重得到新的数组
	names := lo.Uniq[string]([]string{"Samuel", "Marc", "Samuel"})
	// []string{"Samuel", "Marc"}
	fmt.Println(names)

	// 将一个数组进行map运算得到一个新的其他类型的数组
	mapValues := lo.Map[int64, string]([]int64{1, 2, 3, 4}, func(x int64, _ int) string {
		return strconv.FormatInt(x, 10)
	})
	// []string{"1", "2", "3", "4"}
	fmt.Println(mapValues)
}
