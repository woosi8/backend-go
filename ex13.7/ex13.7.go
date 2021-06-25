package main

import (
	"fmt"
	"unsafe"
)



type User struct {
	A int8
	C int8
	E int8 //1바이트 (1바이트를 앞으로 모아놓으니 패딩이 줄어들수 밖에)
	B int  //8바이트
	D int
	// A int8
	// B int
	// C int8
	// D int
	// E int8
}


func main()  {
	user := User{1,2,3,4,5}
	fmt.Println(unsafe.Sizeof(user))
}


