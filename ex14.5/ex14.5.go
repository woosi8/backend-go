package main

import (
	"fmt"
)

type User struct{
	Name string
	Age int

}

func NewUser(name string,age int) *User {
	var u = User{name,age}
	return &u //stack이 아닌 heap(전역변수)로 가게 된다
}

func main()  {
	userPointer := NewUser("AAA",23)
	fmt.Println(userPointer)
}
