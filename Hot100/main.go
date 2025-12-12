package main

/**
本节练习leecode_Hot100里的经典题
1.字母异位词分组
2.最长连续序列
*/
import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(longestConsecutive([]int{0, 1, 2, 4, 8, 5, 6, 7, 9, 3, 55, 88, 77, 99, 999999999}))
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
