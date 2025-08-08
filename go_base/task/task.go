package go_base

import (
	"fmt"
	"sort"
)

func SingleNumber(nums []int) int {
	countMap := make(map[int]int)
	for _, num := range nums {

		countMap[num]++
	}
	for num, count := range countMap {
		if count == 1 {
			return num
		}
	}
	return -1
}
func IsPalindrome(x int) bool {
	s := fmt.Sprintf("%d", x)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func IsValid(s string) bool {
	stack := []rune{}
	mapping := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, v := range s {
		if matchLeft, ok := mapping[v]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != matchLeft {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}
	return len(stack) == 0
}
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[0][i] != strs[j][i] {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}
func TwoSum(nums []int, target int) []int {
	lookup := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if j, ok := lookup[complement]; ok {
			return []int{j, i}
		}
		lookup[num] = i
	}

	return nil // 按题意不会到这一步
}

func Merge(intervals [][]int) [][]int {

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{intervals[0]}
	for i := 0; i < len(intervals); i++ {
		last := res[len(res)-1]
		curr := intervals[i]

		if curr[0] <= last[1] {
			last[1] = max(last[1], curr[1])
		} else {
			res = append(res, curr)
		}
	}
	return res
}
func RemoveDuplicates(nums []int) int {
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}
func PlusOne(digits []int) []int {
	for n := len(digits) - 1; n >= 0; n-- {
		digits[n]++
		if digits[n] < 10 {
			return digits
		}
		digits[n] = 0
	}
	return append([]int{1}, digits...)
}
