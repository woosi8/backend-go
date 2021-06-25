package main

import (
	"fmt"
)

type User struct{
	Name string
	ID string
	Age int
}

type VIPUser struct{
	// UserInfo User
	User  // 이름을 안주고 바로 하는 embeded  내장된 필드
	VIPLevel int
	Price int
	Name string // name중복시 여기가 우선순위를 갖는다
}


func main()  {
	user := User{"송하나","hana",23}
	vip := VIPUser{
		User{"화랑","hwarang",48},
		3,
		250,
		"아무개",
	}

	fmt.Printf("유저: %s ID: %s 나이: %d\n", user.Name,user.ID,user.Age)
	fmt.Printf("VIP유저: %s ID: %s 나이: %d VIP레벨:%d VIP가격: %d만원\n", 
				// vip.UserInfo.Name,vip.UserInfo.ID,vip.UserInfo.Age, vip.VIPLevel,vip.Price)
				// vip.Name,
				vip.User.Name, //필드명이 겹칠경우 원하는 곳에 접근할때 User는 field명이 없으니깐
				vip.ID,vip.Age, vip.VIPLevel,vip.Price)

}