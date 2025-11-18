package main

import "fmt"

/*
 包含常用的几个查找算法
-1. 二分查找
*/

func main() {

	// 二分查找测试
	fmt.Println(binarySearch(5, []int{1, 3, 5, 7, 9, 11, 13}))
}

func binarySearch(k int, array []int) int {
	n := len(array)
	if n == 0 {
		return -1
	}
	mid := n / 2
	if array[mid] == k {
		return mid
	}
	if mid-1 >= 0 {
		left := binarySearch(k, array[:mid])
		if left != -1 {
			return left
		}
	}
	if mid+1 < n {
		right := binarySearch(k, array[mid+1:])
		if right != -1 {
			return mid + right + 1
		}
	}
	return -1
}
