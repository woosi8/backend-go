package main

import "fmt"

func main()  {
	var a int = 10
	var b int = 20
	var f float64 = 32799438743.8297
	fmt.Print("a:", a,"b:",b)
	fmt.Println("a:", a,"b:",b, "f:", f) // ln은 빈칸이 하나하나에 생긴다
	fmt.Printf("a:%d b: %d f:%f\n", a,b,f) // %decimal 정수, 위에서는 f가 실수가 지수로 출력이 되었는데 여기서 %f는 실수값으로 표현된다

}

//go mod init  goprojects/hello
// go build
// hello.exe

