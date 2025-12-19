package main

/**
本节练习leecode_Hot100里的经典题
1.字母异位词分组
2.最长连续序列
3.乘最多水的容器_ac（双指针）
4.三数和为零
*/
import (
	"fmt"
	"slices"
)

func main() {
	//fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	//fmt.Println(longestConsecutive([]int{0, 1, 2, 4, 8, 5, 6, 7, 9, 3, 55, 88, 77, 99, 999999999}))
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))

}

/*
1.字母异位词分组 - 将相同字母组成的不同单归为一组，共多组返回
解题思路：如果两个字符串从小到大排序后相等，那么两个字符串就互为字母异位词，否则不是。
*/
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		s := sortString(strs[i])
		if m[s] == nil {
			m[s] = []string{}
			m[s] = append(m[s], strs[i])
		} else {
			m[s] = append(m[s], strs[i])
		}
	}
	result := [][]string{}
	for _, v := range m {
		result = append(result, v)
	}
	return result
}
func sortString(s string) string {
	ints := make([]string, len(s))
	for i := 0; i < len(s); i++ {
		ints[i] = string(s[i])
	}
	slices.Sort(ints)
	result := ""
	for _, s2 := range ints {
		result += s2
	}
	return result
}

/*
2.最长连续序列
这是本题的标准 O(N) 解法。
思路：1.去重与查询： 将所有数字放入一个 Map 中。

	2.只找“开头”： 遍历 Map 中的数字。如果一个数字 x，在 Map 中存在 x-1，说明 x 不是连续序列的起点，直接跳过（因为我们会在处理 x-1 或更前面的数时计算到 x）。
	3.计算长度： 如果 x-1 不存在，说明 x 是起点。通过 while 循环检查 x+1, x+2 是否存在，直到断开。
*/
func longestConsecutive(nums []int) int {
	has := make(map[int]bool, len(nums))

	for i := 0; i < len(nums); i++ {
		has[nums[i]] = true
	}

	var result = 0
	for x := range has {
		if has[x-1] { //如果它不是起点
			continue
		}

		y := x + 1
		for has[y] { //遍历以x开头连续的所有元素，并以y+1结尾
			y++
		}
		if y-x > result {
			result = y - x
		}
	}
	return result

}

/*
3.乘最多水的容器 func maxArea(height []int) int;
解题思路：暴力或双指针（优先移动短板指针） _一次ac
*/

/*
4.三数之和为零
思路：
1.对数组进行排序。
2.遍历排序后数组：
若 nums[i]>0：因为已经排序好，所以后面不可能有三个数加和等于 0，直接返回结果。
对于重复元素：跳过，避免出现重复解
令左指针 L=i+1，右指针 R=n−1，当 L<R 时，执行循环：
        当nums[i]+nums[L]+nums[R]==0，执行循环，判断左界和右界是否和下一位置重复，去除重复解。并同时将 L,R 移到下一位置，寻找新的解
        若和大于 0，说明 nums[R] 太大，R 左移
        若和小于 0，说明 nums[L] 太小，L 右移
* 主要记住固定元素i，按和与0的大小比较决定双指针移向,麻烦点在于去重_一个是固定点i的去重;一个是组合l_r的去重
*/

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	fmt.Println(nums)
	var result = [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		// 去重i_重复
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}

		var l = i + 1
		var r = len(nums) - 1

		for l < r {
			if nums[l]+nums[r]+nums[i] == 0 {
				//只有不跳过时增加

				temp := []int{}
				temp = append(temp, nums[l], nums[i], nums[r])
				result = append(result, temp)

				// 去重组合l_r
				valL := nums[l]
				valR := nums[r]
				for l < r && nums[l] == valL {
					l++
				}
				for l < r && nums[r] == valR {
					r--
				}

			} else if nums[l]+nums[r]+nums[i] > 0 {
				r--
			} else {

				l++
			}
		}

	}
	return result
}
