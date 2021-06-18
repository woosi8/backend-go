package main

import "fmt"

func main()  {
	// 실수연산은 오차를 조심해야 한다
	var a float32 = 1234.523 //float32는 7자리까지만 나온다
	var b float32 = 3456.123
	var c float32 = a * b //float64는 15자리까지다
	var d float32 = c * 3

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}