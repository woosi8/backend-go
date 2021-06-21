package main

import "fmt"


func main()  {
	var a float64 = 0.1 //8바이트 실수
	var b float64 = 0.2
	var c float64 = 0.3
	fmt.Printf("%f + %f == %f : %v\n", a, b, c, a+b == c)
	fmt.Println(a+b)
}