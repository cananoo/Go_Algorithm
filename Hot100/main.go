package main

/**
本节练习leecode_Hot100里的经典题
1.字母异位词分组
*/
import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
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
