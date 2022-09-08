package cheetsheet

import "sort"

/*
   双指针
使用条件
1. 滑动窗口 (90%)
2. 时间复杂度要求 O(n) (80%是双指针)
3. 要求原地操作，只可以使用交换，不能使用额外空间 (80%)
4. 有子数组 subarray /子字符串 substring 的关键词 (50%)
5. 有回文 Palindrome 关键词(50%)

时间复杂度:O(n)  空间复杂度:O(1)
*/

// 种类1
// 对撞指针：以双指针为左右边界（也就是「数组」的左右边界）计算出结果（不一定是排序后的数组，要看具体题目）
// 典型例题：leetcode11 盛最多水的容器, leetcode15 三数之和

// leetcode11题解（双指针模板）：
func maxArea(height []int) int {
	var res int
	for i, j := 0, len(height)-1; i < j; {
		if height[i] < height[j] {
			res = max(res, height[i]*(j-i))
			i++
		} else {
			res = max(res, height[j]*(j-i))
			j--
		}
	}
	return res
}

// 模板归纳
/*
  1     2     3     4     5     6
 left                         right
*************
for 循环数组：
如果 满足条件1 则 left++
否则 right++
在上边的这个流程里计算res
*************
盛水容器这道题很巧妙的利用了蓄水的物理特性，可以通过对撞指针的方法算出具体的值
*/

// 再如 leetcode15 三数之和
// 三数之和这道题是排序后的数组，通过排序后的有序性计算对应的值
// 由于题目的特殊性，增加了一些过滤重复数字的额外逻辑
// 但主体仍然是：
/**********
for 循环数组：
如果 满足条件1 则 left++
否则 right++
在上边的这个流程里计算res
***********/

func threeSum(nums []int) (ans [][]int) {
	n := len(nums)
	sort.Ints(nums)

	// 枚举 a
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// c 对应的指针初始指向数组的最右端
		right := n - 1
		target := -1 * nums[first] // 计算target
		// 枚举 b
		for left := first + 1; left < right; {

			if left > first+1 && nums[left] == nums[left-1] {
				left++
				continue
			}
			// 需要和上一次枚举的数不相同(过滤条件)
			if right < n-1 && nums[right] == nums[right+1] {
				right--
				continue
			}
			// 开始对撞指针
			if nums[left]+nums[right] > target {
				right--
			} else if nums[left]+nums[right] < target {
				left++
			} else {
				ans = append(ans, []int{nums[first], nums[left], nums[right]})
				left++
				right-- // 不会有两个相同数字的threesum
				continue
			}
		}
	}
	return ans
}

/**************************************************************************/

// 种类2
// 快慢指针：两个指针从同一侧开始遍历数组，将这两个指针分别定义为快指针（fast）和慢指针（slow），
// 两个指针以不同的策略移动，直到两个指针的值相等（或其他特殊条件）为止
// 典型例题：leetcode26 删除有序数组的重复项, leetcode19 删除链表中倒数第N个节点

// 模板题：leetcode26
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

/* 模板归纳
******************************
// 在数组中，同向的快慢指针一般是：
// 赋值初始化slow, fast := val, val
// for循环快指针：
// 循环中，if 满足条件1
// 操作业务代码，慢指针跟上
// 循环结束
******************************
*/

// 再如leetcode19 删除链表中倒数第N个节点，是从链表操作运用快慢指针，链表是更容易操作快慢指针的数据结构
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	fast, slow := head, dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for ; fast != nil; fast = fast.Next {
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

/* 模板归纳
******************************
// 在链表中，同向的快慢指针一般是：
// 赋值初始化slow, fast := val, val
// for循环操作快指针，让快指针先找到对应的节点
// 或者在for循环中对待快慢指针的操作不同
// 循环结束后，同时操作再快慢指针
// 找到相应节点/达成条件后结束
******************************
*/
