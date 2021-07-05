package main

import (
	"fmt"
	"math/rand" //math 안에 있는  rand라는 뜻(사용할때는 마지막 rand만 쓰면된다)
	_ "strings" //패키지 초기화에 따른 부가효과때문에 쓰지 않더라도 있어야 할때
)


func main()  {
	fmt.Println(rand.Int())
}