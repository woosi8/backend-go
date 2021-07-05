package main

import (
	"fmt"
	"strings"
)

func ToUpper1(str string) string  {
	var rst string
	for _, c := range str{
		if c >= 'a' && c <= 'z'{ // 소문자일경우라는 뜻
			rst += string('A' + (c-'a')) //소문잘르 대문자로 바꿔준다
		} else{
			rst += string(c)
		}
	}
	return rst
}


//위에 합산 방법은 매번 새로운 공간이 할당되서 좋지않다. 
// builder를 사용하면 복사하는게 아니라 공간을 늘려주는(새로운 공간이 매번 할당) 개념
func ToUpper2(str string) string  {
	var builder strings.Builder
	for _, c := range str{
		if c >= 'a' && c <= 'z'{ // 소문자일경우라는 뜻
			builder.WriteRune('A' + (c - 'a'))
		} else{
			builder.WriteRune(c)
		}
	}
	return builder.String()
}

// func ToUpper3() {
// 	fmt.Println(strings.ToUpper("Gopher"))
// }
func main()  {
	var str string = "Hello World"

	fmt.Printf(ToUpper1((str)))
	fmt.Printf(ToUpper2((str)))
}