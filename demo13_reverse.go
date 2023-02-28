package main

import (
	"fmt"
	"math"
	"strconv"
)

func reverse(x int) int {
	if x == 0 {
		return x

	} else if x < 0 {
		fl := math.Abs(float64(x))
		x = int(fl)
		s1 := reverse_string(strconv.Itoa(x))
		re_num, _ := strconv.Atoi(s1)
		return -(check(re_num))
	} else {
		s2 := reverse_string(strconv.Itoa(x))
		re_num, _ := strconv.Atoi(s2)
		return check(re_num)
	}

}
func check(a int) int {
	if a < -(1<<31) || a > 1<<31-1 {
		return 0
	} else {
		return a
	}
}

func reverse_string(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]

	}
	return string(runes)
}
func main() {
	x := 123
	x = reverse(x)
	fmt.Println(x)
}

func good_reverse(x int) int {
	tmp := 0
	for x != 0 {
		tmp = tmp*10 + x%10
		x = x / 10
	}
	if tmp > 1<<31-1 || tmp < -(1<<31) {
		return 0
	}
	return tmp
}
