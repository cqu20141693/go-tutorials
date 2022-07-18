package leetcode

/*
	两数之和
*/
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for index, num := range nums {
		if v, ok := m[target-num]; ok {
			return []int{v, index}
		} else {
			m[num] = index
		}
	}
	return []int{}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	go 对象创建{}
	通过& 获取对象的地址引用，对象传递时值拷贝，引用传递是地址，修改可见
	go 中只存在for 循环，没有while循环
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var curr *ListNode
	carry := 0
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else {
		for l1 != nil || l2 != nil || carry > 0 {
			node := &ListNode{}
			carryVal := 0
			if l1 == nil {
				if l2 == nil {
					carryVal = carry
				} else {
					carryVal = l2.Val + carry
				}
			} else if l2 == nil {
				carryVal = l1.Val + carry
			} else {
				carryVal = l1.Val + l2.Val + carry
			}

			if carryVal > 9 {
				node.Val = carryVal - 10
				carry = 1
			} else {
				node.Val = carryVal
				carry = 0
			}

			if l1 != nil {
				l1 = l1.Next
			}
			if l2 != nil {
				l2 = l2.Next
			}
			if head == nil {
				head = node
				curr = node
			} else {
				curr.Next = node
				curr = node
			}

		}
		return head
	}

}
