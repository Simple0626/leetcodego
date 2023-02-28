package main

import (
	"fmt"
	"strings"
)

func longestCommonprefix(strs []string) string {
	var com string
	if len(strs) == 0 {
		return ""
	}
	com = strs[0]
	for i := 0; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], com) {
			com = strs[0][0 : len(com)-1]
			if com == "" {
				return ""
			}
		}
	}
	return com

}

func main() {
	var strs = []string{"flower", "flow", "flight"}
	var str string

	str = longestCommonprefix(strs)
	fmt.Println(str)

}
