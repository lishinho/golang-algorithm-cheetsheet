package cheetsheet

import "sort"

/*
  排序算法sorting
时间复杂度:
○ 快速排序(期望复杂度) : O(nlogn)
○ 归并排序(最坏复杂度) : O(nlogn)
空间复杂度:
○ 快速排序 : O(1)
○ 归并排序 : O(n)

golang有自己封装好的排序api，不过这里的排序(快速排序和归并排序)思想(分治，快排)也会在很多题中刷到。
*/

// 封装好的api
// IntSlice Float64Slice StringSlice 一般封装了这三种slice接口，分别对应各类型切片
func searchSort() {
	sort.Ints([]int{1, 3, 2})                   //整数数组排序，顺序是从小到大
	sort.IntsAreSorted([]int{1, 2, 3})          // 返回值bool，判断输入整数数组是否有序
	sort.Reverse(sort.IntSlice([]int{1, 2, 3})) // 输出输入的order倒序

	sort.Strings([]string{"aaa", "bbb", "ccc"}) // string也可以排序，按照ascii字符大小

	sort.SearchInts([]int{1, 3, 2, 3}, 3)                            // 在数组中找到target的位置，相同的元素出现返回leftmost
	sort.SearchStrings([]string{"aaaa", "ccc", "bbb", "ccc"}, "ccc") // 字符串也可以通过sort包提供的api找到

	// sort.Slice(slice interface{}, less func(i, j int) bool)
	// 提供了自己定义的结构体的排序方法，排序规则可以自定义

	// sort.Stable(data Interface) 方法 提供了稳定性排序方法
	// 稳定性排序即如果数据的键值相同，在排序后相对位置与排序前相同，即稳定排序。
}

//**************************************************
//基础部分
/*
  快速排序实现（以下例子均通过[]int实现）
*/
// 双指针在无序数组区间内排序
func partition(list []int, low, high int) int {
	pivot := list[low] //导致 low 位置值为空
	for low < high {
		//high指针值 >= pivot high指针左移
		for low < high && pivot <= list[high] {
			high--
		}
		//填补low位置空值
		//high指针值 < pivot high值 移到low位置
		//high 位置值空
		list[low] = list[high]
		//low指针值 <= pivot low指针右移
		for low < high && pivot >= list[low] {
			low++
		}
		//填补high位置空值
		//low指针值 > pivot low值 移到high位置
		//low位置值空
		list[high] = list[low]
	}
	//pivot 填补 low位置的空值
	list[low] = pivot
	return low
}

func QuickSort(list []int, low, high int) {
	if high > low {
		//位置划分
		pivot := partition(list, low, high)
		//左边部分排序
		QuickSort(list, low, pivot-1)
		//右边排序
		QuickSort(list, pivot+1, high)
	}
}

// 归并排序
func mergeSort(s []int) []int {
	n := len(s)
	if n == 1 {
		return s //最后切割只剩下一个元素
	}
	m := n / 2
	leftS := mergeSort(s[:m])
	rightS := mergeSort(s[m:])
	return merge(leftS, rightS)
}

//把两个有序切片合并成一个有序切片
func merge(l []int, r []int) []int {
	lLen := len(l)
	rLen := len(r)
	res := make([]int, 0)

	lIndex, rIndex := 0, 0 //两个切片的下标，插入一个数据，下标加一
	for lIndex < lLen && rIndex < rLen {
		if l[lIndex] > r[rIndex] {
			res = append(res, r[rIndex])
			rIndex++
		} else {
			res = append(res, l[lIndex])
			lIndex++
		}
	}
	if lIndex < lLen { //左边的还有剩余元素
		res = append(res, l[lIndex:]...)
	}
	if rIndex < rLen {
		res = append(res, r[rIndex:]...)
	}

	return res
}

//***************************************************

// 例题部分
//***************************************************
// 例题：leetcode-56合并区间  leetcode-148排序链表
// leetcode56:很多情况下，结构体的sort可以拆分成不同基础元素的sort，再做拼合
// 需要敏锐的使用sort做排序，尽量避免自己写一个slice的接口使用sort.Slice，但极端情况下也可以用
func mergeLeetcode56(intervals [][]int) (res [][]int) {
	n := len(intervals)
	var (
		starts []int
		ends   []int
	)
	for _, interval := range intervals {
		starts = append(starts, interval[0])
		ends = append(ends, interval[1])
	}
	if len(starts) == 0 || len(ends) == 0 {
		return nil
	}
	sort.Ints(starts)
	sort.Ints(ends)
	for i, j := 0, 0; i < n; i++ {
		if i == n-1 || starts[i+1] > ends[i] { //中间出现了不连续，记录前一个连续值
			res = append(res, append([]int(nil), starts[j], ends[i]))
			j = i + 1
		}
	}
	return res
}

// leetcode-148排序链表
// 充分利用分治思想：终点是如何把两个有序切片/链表合并成一个有序切片/链表，即merge部分
//
func mergeNode(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	temp, temp1, temp2 := dummyHead, head1, head2
	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}
	if temp1 != nil {
		temp.Next = temp1
	} else if temp2 != nil {
		temp.Next = temp2
	}
	return dummyHead.Next
}

func sortNode(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == tail {
		head.Next = nil
		return head
	}

	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}

	mid := slow
	return mergeNode(sortNode(head, mid), sortNode(mid, tail))
}

func sortList(head *ListNode) *ListNode {
	return sortNode(head, nil)
}

//***************************************************
