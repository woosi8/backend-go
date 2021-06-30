package main

import (
	"fmt"
)

func main()  {
	str := "Hello 월드"
	// rang로 한글자씩 순회. 왜냐면 len은 비트의 길이를 가져오기때문에 못쓴다
	for _, v:= range str {
		fmt.Printf("타입:%T 값:%d 문자값:%c\n", v,v,v)
	}
}