package leetcode

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	// 结构体定义
	type args struct {
		nums   []int
		target int
	}
	// 结构体初始化{}，切片初始化{}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"两数之和",
			args{[]int{2, 5, 6, 7, 11}, 9},
			[]int{0, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"Test_addTwoNumbers ",
			args{&ListNode{1, nil},
				&ListNode{1, nil}},
			&ListNode{2, nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
