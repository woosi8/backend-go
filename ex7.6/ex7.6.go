package main

import (
	"fmt"
)

func printNo(n int)  {
	if n ==0 {
		return //return 호출된 자리로 돌아간다
	}
	fmt.Println(n)
	printNo(n-1) // 여기로
	fmt.Println("After", n)
}


func main()  {
	printNo(3)
	
}

