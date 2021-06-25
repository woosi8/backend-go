package main

import (
	"fmt"
)

func main()  {
	var t [5]float64 = [5]float64{24.0,25.9,27.8,26.9,26.2}
	for i , v:= range t{ //range : 각 요소들을 순회해준다 for문 안에서 (인덱스,요소값 순서로 리턴 한다)
	// for _, v := range t // v값만 리턴하기 위해서 _빈칸 지시자가 필요하다	
		fmt.Println(i,v)
	}
}