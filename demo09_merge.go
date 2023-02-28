package main

import (
	"fmt"
	"sort"
)

func merge(nums1, nums2 []int, n, m int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, nums2, n, m)
	fmt.Printf("%v", nums1)
}
