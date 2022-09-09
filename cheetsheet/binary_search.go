package cheetsheet

import "sort"

/*
二分法

使用条件
1. 排序数组 (30-40%是二分)
2. 当面试官要求你找一个比 O(n) 更小的时间复杂度算法的时候(99%)
3. 找到数组中的一个分割位置，使得左半部分满足某个条件，右半部分不满足(100%)
4. 找到一个最大/最小的值使得某个条件被满足(90%)

时间复杂度:O(logn)  空间复杂度:O(1)
*/

func binarySearch(nums []int, target int) int {
	// a >= target 表示找到第一个相同元素出现的位置，否则找到排序后左侧比target小的那一个
	return search(len(nums), func(i int) bool { return nums[i] >= target })
	// 也可以直接用golang的sort包做
	//return sort.Search(len(nums), func(i int) bool {return nums[i] >= target})
}

// 模板
func search(n int, f func(i int) bool) int {
	m := 0
	for m < n {
		t := int(uint(m+n) >> 1) // 位运算速度更快，uint可以防止溢出
		if !f(t) {
			m = t + 1
		} else {
			n = t
		}
	}
	return m
}

// 例题：leetcode34：在排序数组中查找元素的第一个和最后一个位置
func searchRange(nums []int, target int) []int {
	leftmost := sort.SearchInts(nums, target)
	if leftmost == len(nums) || nums[leftmost] != target {
		return []int{-1, -1}
	}
	rightmost := sort.SearchInts(nums, target+1) - 1
	return []int{leftmost, rightmost}
}

/*
 二分法的一个思想是：在排序好的数组中，快速找到自己想要的数的位置，使用对半比较大小的方式，而不是循环迭代。
 1     2     3     4     5     6
left        mid              right
关键点是优雅地找到这个排序数组的中间数mid := int(uint(left+right) >> 1)
这样就能一半一半地淘汰元素，快速找到自己的target
*/
