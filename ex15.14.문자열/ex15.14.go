package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main()  {
	
	str1 := "Hello 월드"
	str2 := str1

	stringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str1)); // &str1 : string pointer타입 , unsafe.Pointer : 타입변환 raw pointer
	stringHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&str2)) // reflect.StringHeader: structer type ( data,len)


	fmt.Print(stringHeader1)
	fmt.Print(stringHeader2)
	
}