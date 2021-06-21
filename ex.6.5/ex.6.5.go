package main

import (
	"fmt"
)

func main()  {
	var x int8 = 127
	fmt.Printf("%d < %d+1: %v\n", x, x, x < x+1)
	fmt.Printf("x\t = %4d, %08b\n",x,x) //한번은 4자리정수로, 한번은 2진수로
	fmt.Printf("x+1\t = %4d, %08b\n", x+1,uint8(x+1))
}

// 정수 오버플로우
// 가장 큰값에서 +1을 하면 가장 작은값으로 변한다.  -128~127