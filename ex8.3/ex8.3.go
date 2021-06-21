package main

import (
	"fmt"
)

// const Pig int = 0
// const Cow int = 1
// const Chicken int = 2

// const (
// 	Pig int =0
// 	Cow int = 1
// 	Chicken int = 2
// )

const (
	Pig int = iota + 1 //iota : 자동으로 1씩 증가
	Cow 
	Chicken 
)


func PrintAnimal(animal int)  {
	if animal == Pig {
		fmt.Println("oweowe")
	} else if animal == Cow{
		fmt.Println("cowww")
	}  else if animal == Chicken{
		fmt.Println("kokio")
	}  else{
		fmt.Println("...")
	} 
}

func main()  {
	PrintAnimal(Cow)
	PrintAnimal(Pig)
	PrintAnimal(7)
}