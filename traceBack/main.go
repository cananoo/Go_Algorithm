package main

import (
	"fmt"
	"strings"
)

/*
总结和回溯相关的编程应用题 -- 回溯的本质是递归+中间变量
1. 全排列
*/

func main() {
	fullPermutation("abc", "")

}

// 1.全排列
/*
解题：回溯的本质就是每一轮要不要的问题，创建 要/不要的两条路径，选择要就要删去原有的，选择 不要 或 后面要 则要恢复删去的再进入下一轮
 -1.像全排列这种全都要的但要的顺序不一样的 可以用for + 回溯的方法.
 -2.像SubSet/背包 这种组合问题 可以不用for ，直接创建路径选择即可.
 -3.不管是上述哪种问题，都可以构造解空间树-理清思路后再解题。
*/
func fullPermutation(input string, temp string) {
	if len(input) == 0 {
		fmt.Println(temp)
		return
	}
	n := len(input)
	for i := 0; i < n; i++ {
		var later = input
		// 加上该字符
		temp += string(input[i])
		// 将此字符从原串抹去
		s := string(input[i])
		input := strings.Replace(input, s, "", 1)
		fullPermutation(input, temp)

		//恢复
		temp = strings.Replace(temp, string(temp[len(temp)-1]), "", 1)
		input = later
	}

}
