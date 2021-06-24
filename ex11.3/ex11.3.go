package main

import (
	"bufio"
	"fmt"
	"os"
)


func main()  {
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("press enter")
		var number int
		_, err := fmt.Scanln(&number)
		// _,: scanln의 리턴값이 두개가 리턴되는데 첫번째가 읽어온 값의 개수, 두번재가 에러여부, 여기선 읽어온 값의 개수가 필요없어서 _치고 날려버린다
		if err !=nil {
			fmt.Println("put number")
			// '': 문자하나를 나타내고 "" :문자 열을 나타낸다
			//키보드 버퍼를 지웁니다.
			stdin.ReadString('\n') //개행 문자가 나올때까지 키보드 버퍼를 읽어오는데 읽어온 값이 들어온다. 여기선 에러가 나온다
			continue // 후처리로 이동, 후처리 없어서 다시 for문 반복
		}
		fmt.Printf("입력하신 숫자는 %d 입니다.\n", number)
		if number%2 == 0 {
			break //짝수면 멈춰라
		}
	}
	fmt.Println("for문이 종료되었습니다.")
}