package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var (
		n = 0
		m = 0
	)

	for i := 0; i < len(nums); i++ {
		n = i
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				n = i
				m = j
				return []int{n, m}
			}
		}
	}
	return []int{}
}

func main() {
	num := []int{1, 2, 3, 5, 6}
	target := 6
	result := twoSum(num, target)
	fmt.Println(result)
}
