package main

import (
	"fmt"
)

// func main()  {
// 	var a int = 500
// 	var p *int
// 	p = &a

// 	fmt.Printf("p의 값 :%p\n", p)
// 	fmt.Printf("p의 메모리값 :%d\n", *p)
// 	*p = 100
// 	fmt.Printf("a의 값: %d\n", a)

// }
func main()  {
	var a int = 10
	var b int = 20

	var p1 *int = &a
	var p2 *int = &a
	var p3 *int = &b

	fmt.Printf("p1 == p2 : %v\n", p1 == p2)
	fmt.Printf("p1 == p3 : %v\n", p2 == p3)

}