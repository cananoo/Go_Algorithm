package main

import (
	"fmt"
)

/*
 包含常用的几个排序算法
-1. 快速排序
-2. 归并排序
*/

func main() {
	// 1. 测试快排
	fmt.Printf("%v\n", quickSort(7, []int{3, 6, 8, 10, 1, 2, 1}))

	// 2. 测试归并排序
	fmt.Println(merge([]int{3, 27, 38, 43}, []int{9, 10, 82})) //测试归并算法是否成功
	fmt.Println(mergeSort([]int{38, 27, 43, 3, 9, 82, 10}))

}

/*
1）选择基准元素：从列表中选择一个元素作为基准（pivot）。这里选择方式可以是第一个元素
2）分区：将列表重新排列，使得所有小于基准元素的元素都在基准的左侧，所有大于基准元素的元素都在基准的右侧。基准元素的位置在分区完成后确定。
3）递归排序：对基准元素左侧和右侧的子列表分别递归地进行快速排序。
4）合并：由于分区操作是原地进行的，递归结束后整个列表已经有序。
最好时间复杂度：O(n*logn) 最坏时间复杂度：O(n^2)-当取首为基准且为正序或逆序时-分区极度不平衡-递归深度 O (n)  稳定性：不稳定

* 进阶-三数取中法优化：最坏O(n*logn)-大幅降低选到极值的概率，保证分区相对平衡
https://www.cnblogs.com/chengxiao/p/6262208.html
*/
func quickSort(n int, array []int) []int {
	if n <= 1 {
		return array
	}
	pivot := array[0]
	var left []int
	var right []int
	for i := 1; i < n; i++ {
		if array[i] < pivot {
			left = append(left, array[i])
		} else {
			right = append(right, array[i])
		}
	}
	left = quickSort(len(left), left)
	right = quickSort(len(right), right)
	temp := append(left, pivot)
	ints := append(temp, right...)
	return ints
}

/*
1)分解（Divide）：将待排序的数组分成两个子数组，每个子数组包含大约一半的元素。
2)解决（Conquer）：递归地对每个子数组进行排序。
3)合并（Combine）：将两个已排序的子数组合并成一个有序的数组。
前提：有一个能够将两个有序数组合并成一个有序数组的方法
https://www.runoob.com/w3cnote/merge-sort.html
*/
func mergeSort(array []int) []int {
	if len(array) == 0 || len(array) == 1 {
		return array
	}
	mid := len(array) / 2
	left := array[:mid]
	right := array[mid:]
	left = mergeSort(left)
	right = mergeSort(right)
	sort := merge(left, right)
	return sort
}

func merge(left, right []int) []int {
	var result []int
	var i, j = 0, 0 //两个数组的滑动指针
	for {
		if i < len(left) && j < len(right) && left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
		if i >= len(left) || j >= len(right) {
			break
		}
	}
	for {
		if i < len(left) {
			result = append(result, left[i])
		} else {
			break
		}
		i++
	}
	for {
		if j < len(right) {
			result = append(result, right[j])
		} else {
			break
		}
		j++
	}
	return result
}
