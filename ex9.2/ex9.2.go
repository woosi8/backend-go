package main

import (
	"fmt"
)

func main()  {
	temp := 33
	if temp >28 {
		fmt.Println("turn on air")
	} else if temp <= 3 {
		fmt.Println("turn on hitter")
	} else if temp <=18 {
		fmt.Println("go out!")
	} else {
		fmt.Println("hot")
	}

}