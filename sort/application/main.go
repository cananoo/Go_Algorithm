package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
*
总结和排序相关的编程应用题
1.基数日期排序
*/
func main() {

	// 1.基数日期排序测试
	fmt.Println(dateSortByredix(
		[]string{
			"2025-10-01",
			"2025-09-15",
			"2024-12-31",
			"2023-01-01",
		},
	))
}

// 1.基数排序（Radix Sort）实现日期排序 - 日期格式-"2025-09-15"
func dateSortByredix(dates []string) []string {
	// 棑年
	for i := 1; i < len(dates); i++ {
		atoi, _ := strconv.Atoi(strings.Split(dates[i], "-")[0])
		for j := i - 1; j >= 0; j-- {
			s := j + 1
			atoj, _ := strconv.Atoi(strings.Split(dates[j], "-")[0])
			if atoj <= atoi {
				break
			}
			dates[s], dates[j] = dates[j], dates[s]
		}
	}
	// 棑月
	for i := 1; i < len(dates); i++ {
		atoi, _ := strconv.Atoi(strings.Split(dates[i], "-")[1])
		for j := i - 1; j >= 0; j-- {
			s := j + 1
			yeari, _ := strconv.Atoi(strings.Split(dates[i], "-")[0])
			yearj, _ := strconv.Atoi(strings.Split(dates[j], "-")[0])
			if yeari == yearj { //保证年相等排序月
				atoj, _ := strconv.Atoi(strings.Split(dates[j], "-")[1])
				if atoj <= atoi {
					break
				}
				dates[s], dates[j] = dates[j], dates[s]
			}
		}
	}

	// 棑日
	for i := 1; i < len(dates); i++ {
		atoi, _ := strconv.Atoi(strings.Split(dates[i], "-")[2])
		for j := i - 1; j >= 0; j-- {
			s := j + 1
			yeari, _ := strconv.Atoi(strings.Split(dates[i], "-")[0])
			yearj, _ := strconv.Atoi(strings.Split(dates[j], "-")[0])
			monthi, _ := strconv.Atoi(strings.Split(dates[i], "-")[1])
			monthj, _ := strconv.Atoi(strings.Split(dates[j], "-")[1])
			if yeari == yearj && monthi == monthj { //保证年月相等排序日
				atoj, _ := strconv.Atoi(strings.Split(dates[j], "-")[2])
				if atoj <= atoi {
					break
				}
				dates[s], dates[j] = dates[j], dates[s]
			}
		}
	}

	return dates
}
