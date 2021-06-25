package main

import (
	"fmt"
)

type House struct{
	Address string
	Size int
	Price float64
	Category string

}
func main()  {
	var house House
	house.Address = "seoul gangnam"
	house.Size = 28
	house.Price = 10
	house.Category = "APT"

	// fmt.Println("%v\n", house) //%v는 안에 중가로가 묶인다
	fmt.Printf("주소:%s 사이즈:%d평 가격:%v억원 종류:%s\n", house.Address,house.Size,house.Price,house.Category) //%v는 안에 중가로가 묶인다
}