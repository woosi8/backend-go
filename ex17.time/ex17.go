package main

import (
	"fmt"
	"time"
)

func main()  {

	loc, _ := time.LoadLocation("Asia/Seoul") //리턴값이 2개인데 뒤에거가 필요없이 앞에만 필요할때 ,_ :=
	const longForm = "Jan 2, 2006 at 3:04pm" //포멧
	//ParseInLocation: 문자열을 통해서 시간을 알아온다
	t1, _ := time.ParseInLocation(longForm, "Jun 14,2021 at 10:00pm" , loc) //longForm : 이런 형식이다 라는정의
		fmt.Println(t1,t1.Location(),t1.UTC())

		const shortForm = "2006-Jan-02"
		t2, _ := time.Parse(shortForm, "2021-Jun-14") //그냥 parse는 utc기준
		fmt.Println(t2, t2.Location())
		
		t3, err := time.Parse("2021-06-01 15:20:21", "2021-06-14 20:04:05")
		if err != nil{
			fmt.Println(err)
		}

		fmt.Println(t3, t3.Location())

		d := t2.Sub(t1)
		fmt.Println(d)
}