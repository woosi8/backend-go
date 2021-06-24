package main

import (
	"fmt"
)

func main()  {
	a := 4

	switch a{
	case 1:
		fmt.Println("a == 1")
	case 2 :
		fmt.Println("a == 2")
	case 3,4 : //,콤마는 or
		fmt.Println("a == 3")
	default:
		fmt.Println("a 1= 1,2,3, a")
	}
}