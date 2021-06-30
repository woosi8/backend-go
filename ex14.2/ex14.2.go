package main

import (
	"fmt"
)

type Data struct{
	value int
	data [200]int
}


// func ChangeData(arg Data)  { // arg Data 함수의 인자는 우변(값)
func ChangeData(arg *Data)  { // pointer - data타입의 주소를 값으로 받는다
	arg.value = 999 //pointer변수의 점을 찍으면 (*arg).value = 999 (pointer의 공간이 바뀐다)
	arg.data[100] = 999
}


func main()  {
	var data Data
	// ChangeData(data) //data와 Changedate의 arg 인자는 서로 다른 공간을 갖는다
	ChangeData(&data) //data의 주소값

	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])
}