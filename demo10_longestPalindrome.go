package main

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	res, dp := "", make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	//设置好二维切片的大小
	for i := len(s); i >= 0; i-- {
		for j := i; j < len(s); j++ {
			dp[i][j] = (s[i] == s[j]) && ((j-i < 3) || dp[i+1][j-1])
			if dp[i][j] && (res == "" || j-i+1 > len(res)) {
				res = s[i : j+1]
			}
		}
	}
	return res
}

func main() {
	s := "babad"
	str := longestPalindrome(s)
	println(str)

}
