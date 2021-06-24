package main

import (
	"fmt"
)

func main()  {
	a := 1
	b := 1

OuterFor:    // label 레이블  (사실 잘 안쓰인다 - 강제성이 있어서 내부  stack이 꼬일 위험이 있어서)
	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++{
			if a *b == 45 {
				break OuterFor //OuterFor
			}
		}
	} //여기서 빠져나온다
	fmt.Printf("%d * %d = %d\n", a,b, a*b)
}