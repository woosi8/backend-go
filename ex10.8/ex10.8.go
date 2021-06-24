package main

import (
	"fmt"
)


type ColorType int // 별칭타입 : 새로운 타입 만들기 , 타입 선언, 

const (
	Red ColorType = iota // ColorType 컬러로 사용되기대문에 별칭을 써서 의미를 명확하게 나타내는 코드값이라는것을 나타내기 위해 이렇게 해준다
	Blue
	Green
	Yellow
)

func colorToString(color ColorType) string{ //colortype의 값을 받아서 string으로 변환
	switch color {
	case Red:
		return "RED"
	// break; // go에서는 break안써도 아래로 진행되지 않는다
	case Blue:
		return "Blue"
		fallthrough // break안하고 밑으로 계속 진행하고 싶을때	, green을 리턴하고 yellow 리턴한다
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	default:
		return "Undefined"

	}
}

func getMyFavoriteColor() ColorType {
	return Red
}

func main(){
	fmt.Println("My favorite color is", colorToString(Green))
	// fmt.Println("My favorite color is", colorToString(1))
	// fmt.Println("My favorite color is", colorToString(getMyFavoriteColor()))
}


