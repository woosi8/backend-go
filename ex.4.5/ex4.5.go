package main

import "fmt"

func main() {
	var a int16 = 3456   // a는 int16타입 = 2바이트 정수
	var b int8 = int8(a) // int16 -> int8

	fmt.Println(a,b)

}

// 타입변환 주의사항
// 큰타입에서 작은타입으로 변환될때 값이 잘리는 경우가 있다
// 실수타입을 정수타입으로 바꿀때 소수점은ㅇ 날아간다.