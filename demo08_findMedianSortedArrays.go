package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	lensum := len1 + len2
	if lensum == 0 {
		return float64(0)
	}
	l, r := 0, 0
	a := make([]int, 0, lensum)
	for l < len1 && r < len2 {
		if nums1[l] < nums2[r] {
			a = append(a, nums1[l])
			l++
		} else {
			a = append(a, nums2[r])
			r++
		}
	}
	//将长数组后面的数加到a的后面
	a = append(a, nums1[l:]...)
	a = append(a, nums2[r:]...)
	if lensum%2 != 0 {
		return float64(a[lensum/2])
	} else {
		return (float64(a[lensum/2-1]) + float64(a[lensum/2])) / 2
	}
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	mid := findMedianSortedArrays(nums1, nums2)
	fmt.Printf("中位数是：%f", mid)
}
