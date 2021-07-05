package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main()  {
	var str string = "Hello World"
	var slice []byte = []byte(str) //타입변환할때는 복사하는것 (원본 문자열은 보호함, 다른공간을 만듬)
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&str)) //내부구조
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	fmt.Printf("str:\t%x\n", stringHeader.Data)
	fmt.Printf("slice:\t%x\n", sliceHeader.Data)
}