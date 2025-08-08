package main

import (
	"fmt"

	task "github.com/mygithub-lxh/go_start/go_base/task"
)

func main() {

	// SingleNumber
	fmt.Println("SingleNumber: 找出只出现一次的数字")
	fmt.Println(task.SingleNumber([]int{2, 2, 3}))

	// isPalindrome
	fmt.Println("isPalindrome: 判断整数是否回文")
	fmt.Println(task.IsPalindrome(121))

	// isValid
	fmt.Println("isValid: 判断括号字符串是否有效")
	fmt.Println(task.IsValid("()[]{}"))

	// longestCommonPrefix
	fmt.Println("longestCommonPrefix: 查找最长公共前缀")
	fmt.Println(task.LongestCommonPrefix([]string{"flower", "flow", "flight"}))

	// twoSum
	fmt.Println("twoSum: 找出两数之和为目标值的下标")
	fmt.Println(task.TwoSum([]int{2, 7, 11, 15}, 9))

	// merge
	fmt.Println("merge: 合并重叠区间")
	fmt.Println(task.Merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))

	// removeDuplicates
	fmt.Println("removeDuplicates: 删除有序数组中的重复项")
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3}
	length := task.RemoveDuplicates(nums)
	fmt.Println(nums[:length], "长度:", length)

	// plusOne
	fmt.Println("plusOne: 给表示大整数的数组加一")
	fmt.Println(task.PlusOne([]int{9, 9, 9}))

}
