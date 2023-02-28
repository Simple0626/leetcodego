package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x >= 0 && x < 10 {
		return true
	}
	arr := make([]int, 0, 32)
	for x > 0 {
		arr = append(arr, x%10)
		x = x / 10
	}
	length := len(arr)
	for i, j := 0, length-1; i <= j; i, j = i+1, j-1 {
		if arr[i] != arr[j] {
			return false
		}
	}
	return true
}

func main() {
	var x int
	_, err := fmt.Scanf("%d", &x)
	if err != nil {
		return
	}
	if isPalindrome(x) {
		fmt.Println("是回文数")
	} else {
		fmt.Println("不是回文")
	}

}
