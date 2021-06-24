package main

import (
	"fmt"
)

const (
	MasterRoom uint8 = 1 << iota
	LivingRoom //2
	BathRoom //4
	SmallRoom //8
	Garage //16

)

func SetLight(rooms, room uint8) uint8  {
	return rooms | room //or연산자 둘중에 하나만 1이여도 1이다 (or는 rooms가 1이 있는경우에는 합해진다)
}

func ResetLight(rooms, room uint8) uint8  {
	return rooms &^ room // 비트클리어
}

func IsLightOn(rooms, room uint8) bool  {
	return rooms & room == room
		
}

func TurnLights(rooms uint8)  {
	if IsLightOn(rooms, MasterRoom){
		fmt.Println("turn on light in MasterRoom")
	}
	if IsLightOn(rooms, LivingRoom){
		fmt.Println("turn on light in LivingRoom")
	}
	if IsLightOn(rooms, BathRoom){
		fmt.Println("turn on light in BathRoom")
	}
	if IsLightOn(rooms, SmallRoom){
		fmt.Println("turn on light in SmallRoom")
	}
	
}


func main()  {
	var rooms uint8 = 0
	rooms = SetLight(rooms, MasterRoom)
	rooms = SetLight(rooms, LivingRoom)
	rooms = SetLight(rooms, BathRoom)
	rooms = SetLight(rooms, SmallRoom)

	rooms = ResetLight(rooms, SmallRoom) //스몰룸은 턴off한것
	TurnLights(rooms)
}