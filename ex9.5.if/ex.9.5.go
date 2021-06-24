package main

import (
	"fmt"
)

func HasRichFriend() bool  {
	return true
	
}

func GetFriendscount() int {
	return 3
	
}

func main()  {
	price := 35000

	if price >= 50000 {
		if HasRichFriend(){
			fmt.Println("tide shoes ")
		} else{
			fmt.Println("DutchPay")
		}
	} else if price >= 30000 {
		if GetFriendscount() > 3 {
			fmt.Println("oh shoerace")
		} else{
			fmt.Println("Dutch pay man")
		}
	} else {
		fmt.Println("It's on me")
	}
}


// if 초기문,조건문

// if filename,(uploadfile함수의 string ), success(uploadfile함수의 boolen):= uploadfile(); success{ //성공하면
// 	fmt.Println("Upload success",filenmae)
//    } else{
// 	fmt.Println("Faile dto upload")
//    }