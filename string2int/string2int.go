package main

import (
	"fmt"
	"strconv"
)

func main() {
	num, s := 12, "12"
	//string转int
	s2i, _ := strconv.Atoi(s)
	fmt.Printf("%t,%d\n", s2i, s2i)
	//	string转int64
	s2int64, _ := strconv.ParseInt(s, 10, 64)
	fmt.Printf("%t,%d\n", s2int64, s2int64)
	//	int转string
	i2s := strconv.Itoa(num)
	fmt.Printf("%t,%s\n", i2s, i2s)
	//	int64转string
	int642s := strconv.FormatInt(int64(num), 10)
	fmt.Printf("%t,%s\n", int642s, int642s)
}
