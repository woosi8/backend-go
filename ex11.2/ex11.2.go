package main

import (
	"fmt"
	"time"
)

func main()  {
	i := 1
	for  {
		time.Sleep(time.Second)
		i++
		if i ==4 {
			continue // 바로 후처리로 건너뛴다 , 즉 생략
		}
		if i ==6 {
			break //6에서 멈춤
		}
		fmt.Println(i)
	}
}