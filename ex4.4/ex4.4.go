package main

import "fmt"

func main() {
	a := 3
	var b float64 = 3.5

	// var c int = b
	var c int = int(b) //타입변환
	d := float64(a) * b // 소수를 정수로 바꾸면 소수점 버림

	var e int64 = 7 //int와 int64는 다르게 본다 고에서는
	f := a * int(e)

	fmt.Println(a,b,c,d,e,f)

	// 타입 사이즈
	// unit unsigned 부호없는 정수
	// unit8 1바이트 0~255
	// unit16 2바이트 0~65535
	// init8 1바이트 부호 있는 정수  -128 ~ 127
	// init16 2바이트 부호 있는 정수  -32768 ~ 32767

}