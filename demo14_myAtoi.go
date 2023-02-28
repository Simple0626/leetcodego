package main

func myAtoi(s string) int {
	maxInt, signAllowed, SpaceAllowed, sign, digits := int64(2<<30), true, true, 1, []int{}
	for _, i := range s {
		if i == ' ' && SpaceAllowed {
			//有空格，且可以出现空格
			continue
		}
		if signAllowed {
			//允许出现正负号
			if i == '+' {
				signAllowed = false
				//	正负符号已经找到了
				SpaceAllowed = false
				//	后面就不要有空格了
				continue
			} else if i == '-' {
				sign = -1
				signAllowed = false
				SpaceAllowed = false
				continue
			}

		}
		if i < '0' || i > '9' {
			break
			//	有其他的字符，就跳出循环了，不用在往下查看了

		}
		SpaceAllowed, signAllowed = false, false
		digits = append(digits, int(i-48))
	}
	var num, place int64
	place, num = 1, 0
	lastLeading0Index := -1
	//去除前置0
	for i, d := range digits {
		if d == 0 {
			lastLeading0Index = i
		} else {
			break
		}
	}
	if lastLeading0Index > -1 {
		digits = digits[lastLeading0Index+1:]
	}
	//判断数字大小是否超多范围
	var rtnMax int64
	if sign > 0 {
		rtnMax = maxInt - 1

	} else {
		rtnMax = maxInt
	}

	digitsCount := len(digits)
	for i := digitsCount - 1; i >= 0; i-- {
		num += int64(digits[i]) * place
		place *= 10
		if digitsCount-i > 10 || num > rtnMax {
			return int(int64(sign) * rtnMax)

		}

	}
	num *= int64(sign)
	return int(num)
}
func main() {
	s := "  0000000000012345678"
	num := myAtoi(s)
	println(num)
}
