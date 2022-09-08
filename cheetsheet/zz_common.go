package cheetsheet

/*
  一些常用的快速函数与数据结构
*/

// 较大值
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// 单向链表
type ListNode struct {
	Val  int
	Next *ListNode
}
