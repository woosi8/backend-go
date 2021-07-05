package main

import "fmt"


func main()  {
	var str string = "Hello World"
	var slice []byte = []byte(str) //문자열은 불변이기 때문에 새로운 주소를 가르키게 만들어 준다.

	slice[2] ='a'

	fmt.Println(str)
	fmt.Printf("%s\n", slice)
	
}