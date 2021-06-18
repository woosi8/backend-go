package main

import "fmt"

func main() {
	var a int = 3
	var b int //초기값을 안주면 기본값 0이 할당된다
	var c = 4 //int 생략
	d := 5
	var e = "Hello"
	f := 3.14
	fmt.Println(a,b,c,d,e,f)
}

//타입변환 : 연산의 각 항목의 타입은 반드시 같아야 한다 (같은 정수여도 바이트(사이즈)가 다른경우도 안된다. go언어 같은경우에. (최강타입언어라서))

