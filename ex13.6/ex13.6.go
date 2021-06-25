package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	Age int32 //4
	Score float64 // 8 바이트 인데 결과는 16? 왜 ? 비트가 딱 떨어지는 값이 아니라서 빈 값들이 생긴다(메모리 패딩):
}

func main()  {
	var a int = 8
	user := User{23,77.2}
	fmt.Println(unsafe.Sizeof(user), unsafe.Sizeof(a)) //메모리 사이즈 반환
}