package main

import (
	"fmt"
)

func main()  {
	str := "Hello 월드"
	arr := []rune(str)
	// rune -> int32 별칭타입
	// [6]int 6개배열, []int > slice : 동적배열
	for i := 0; i < len(arr); i++ {
		fmt.Printf("타입:%T 값:%d 문자값:%c\n", arr[i],arr[i],arr[i])
	}
}