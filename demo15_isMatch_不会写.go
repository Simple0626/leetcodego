package main

import "strings"

func isMatch(s, p string) bool {
	if p == ".*" {
		return true
	} else if strings.Contains(p, ".") {

	} else if strings.Contains(p, "*") {

	} else if strings.Contains(p, ".*") {

	} else if s == p {
		return true

	}
	return false
}

func main() {
	s, p := "aa", "aa"
	flag := isMatch(s, p)
	println(flag)

}
