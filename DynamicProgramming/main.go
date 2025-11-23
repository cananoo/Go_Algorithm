package main

import (
	"fmt"
	"math"
)

/*
1. 0-1背包问题
2. LCS最长公共子序列
*/
func main() {
	fmt.Println(oneZeroBackpack([]int{2, 3, 4}, []int{3, 5, 6}, 3, 6))

	fmt.Println(LCS("hellolll", "lloll"))

}

/*
1. 0 - 1 背包问题
解题：
1.构造前0-n物体在容量0-capacity的状态转移矩阵，矩阵的每一个框的位置i,j代表前i个物体组合满足容量j的最优解，那么最后一个框则代表全局最优解
2.将第一行和第一列都填为0，从(1,1)开始到最后一个框执行状态转移函数

	            ->1)单独装入此物品重量>j,直接继承上一行同列总价值元素
				->2)单独装入此物品重量<=j，当前背包能满足容量，记录[此物品价值加上剩余背包重量能满足的最大价值(即 上一行,j-此物品重量上位置的元素)]-x
	            ->3)记录不装入此物品的最大价值，即[上一行同列元素]-y
	            ->4)比较 x,y,取最大值填入此框
				->5)不断执行状态转移，最后一个框即为全局最优解

https://www.bilibili.com/video/BV1pY4y1J7na

类似问题：可以用来解SubSet Problem
*/
func oneZeroBackpack(weights []int, values []int, n int, capacity int) int {
	//初始化状态转移矩阵
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, capacity+1)
	}
	for i := 0; i < capacity+1; i++ {
		dp[0][i] = 0
	}
	for i := 0; i < n+1; i++ {
		dp[i][0] = 0
	}
	//动态转移
	for i := 1; i < n+1; i++ {
		for j := 1; j < capacity+1; j++ {
			//填写框内数字-最优解
			if weights[i-1] <= j { //装的下
				remain := j - weights[i-1]
				// 比较有这个物品能得到的最优解和没有这个物品能得到的最优解的大小
				pre := dp[i-1][j]
				after := values[i-1] + dp[i-1][remain]
				if after > pre {
					dp[i][j] = after
				} else {
					dp[i][j] = pre
				}
			} else { //装不下，继承上一列即可
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n][capacity]
}

/*
2. LCS最长公共子序列
解题：
 1. 构造dp矩阵 - 长和宽分别为两个字串的长度+1，每个元素(i,j)位置代表着子串1前i个字符与子串2前j个字符的最大公共子序列长度,第一行为0，第二行为0，
 2. 从(1,1)到最后一个元素，若当前位置(i,j)即s1[i-1] == s2[j-1],则当前位置值为左上角长度+1：dp[i][j] = dp[i-1][j-1] + 1
    若不相等，则取该位置正上方和正左方的更大值作为此位置的值。
 3. 得到最后一个元素的值后-即最大公共子序列的长度-n，可以按下面方法进行回溯。
    n次循环-若该位置(i，j)满足 s1[i-1] == s2[j-1]，则记录该符号在前方，若不满足，则让元素位置向上或左转移(取决于哪边数更大，相等说明有
    多个解，任选一侧)，当此循环无效，i--增加一次循环次数，然后重复上述操作.

https://www.bilibili.com/video/BV12GgQzaECe
*/
func LCS(s1 string, s2 string) string {
	l1 := len(s1)
	l2 := len(s2)
	// 构造转移矩阵,并初始化第一行第一列为0,默认为0所以省略此步骤
	var dp = make([][]int, l1+1)
	for i := 0; i < l1+1; i++ {
		dp[i] = make([]int, l2+1)
	}
	// 状态转移
	for i := 1; i < +l1+1; i++ {
		for j := 1; j < l2+1; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i][j-1])))
			}
		}
	}

	// 回溯找字串
	n := dp[l1][l2] //最大公共字串的长度
	var result = ""
	for i := 0; i < n; i++ {
		if s1[l1-1] == s2[l2-1] {
			result = string(s1[l1-1]) + result //相同加哪个字符串都行
			l1--
			l2--
		} else if dp[l1-1][l2] >= dp[l1][l2-1] {
			l1--
			i--
		} else {
			l2--
			i--
		}
	}
	return result
}
