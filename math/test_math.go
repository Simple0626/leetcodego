package main

import (
	"fmt"
	"math"
)

func main() {
	i := -100
	//绝对值
	num1 := math.Abs(float64(i))
	fmt.Printf("%.4f\n", num1)
	//向上取整
	num2 := math.Ceil(5.2)
	fmt.Printf("%.4f\n", num2)
	//向下取整
	num3 := math.Floor(5.8)
	fmt.Printf("%.4f\n", num3)
	//取余数
	num4 := math.Mod(11, 3)
	fmt.Printf("%.4f\n", num4)
	//取整数，舍小数
	num5, _ := math.Modf(12.324)
	fmt.Printf("%.4f\n", num5)
	//平方
	num6 := math.Pow(2, 3)
	fmt.Printf("%.4f\n", num6)
	//10的平方
	num7 := math.Pow10(5)
	fmt.Printf("%.4f\n", num7)
	//开方
	num8 := math.Sqrt(8)
	fmt.Printf("%.4f\n", num8)
	//开立方
	num9 := math.Cbrt(27)
	fmt.Printf("%.4f\n", num9)
	//pi
	num10 := math.Pi
	fmt.Printf("%.4f\n", num10)
	//	返回值的整数部分
	num11 := math.Trunc(-1.2324)
	fmt.Printf("%f\n", num11)
	//	返回最大值
	num12 := math.Max(1, 2)
	fmt.Printf("%f\n", num12)
	//	返回最小值
	num13 := math.Max(1, 2)
	fmt.Printf("%f\n", num13)
	//	返回x-y和0比较并返回其中的较大值
	num14 := math.Dim(1, 2)
	fmt.Printf("%f\n", num14)
	//	求2为底自然对数
	num15 := math.Log2(8)
	fmt.Printf("%f\n", num15)
}
