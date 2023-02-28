package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var ret [][]int
	l := len(nums)
	//不足三个直接返回空
	if l < 3 {
		return ret
	}
	sort.Ints(nums) // 排序

	for k := 0; k < l; k++ {
		//最小的数都比0大,后面大的数加起来也不会等于0,所以跳过
		if nums[k] > 0 {
			break
		}
		//前后两个数,不能一样
		if k > 0 && nums[k] == nums[k-1] { // 去重
			continue
		}
		//从k到数组末尾,找一组sum+nums[k]=0的
		for i, j := k+1, l-1; i < j; {
			if j < l-1 && nums[j] == nums[j+1] { // 去重
				j--
				continue
			}
			sum := nums[k] + nums[i] + nums[j]
			if sum > 0 {
				j--
			} else if sum < 0 {
				i++
			} else {
				ret = append(ret, []int{nums[k], nums[i], nums[j]})
				i++
				j--
			}
		}
	}
	return ret
}

func main() {
	nums := []int{-1, -1, 2, -1, 0, 1, 2, -1, -4}
	ret := threeSum(nums)
	fmt.Println(ret)
}
