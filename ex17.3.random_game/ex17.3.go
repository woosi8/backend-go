package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdin = bufio.NewReader(os.Stdin) // NewReader: input으로부터 값을 읽어올수 있는

func InputIntValue() (int,error)  {
	var n int
	_, err := fmt.Scanln(&n) //&n = 주소 (pointer)를 넘겨줘야 (공간을 넘겨주기위해)값을 채울수 있다, 입력값을 받을때 scanLn으로 받았고
	if err != nil {
		stdin.ReadString('\n') //에러가 있으면 개행문자가 나올때까지 stdin에서 읽어라 (입력버퍼에 문자가 나올때까지 비워주는 역할)
	}
	return n, err

}

func main()  {
	rand.Seed(time.Now().UnixNano())	

	r := rand.Intn(100)
	cnt := 1
	for{
		fmt.Println("숫자를 입력하세요 : ")
		n, err := InputIntValue()
		if err != nil{
			fmt.Println("숫자만 입력하세요")
		} else{
			if n > r {
				fmt.Println("입력하신 숫자가 더 큽니다.")
			}else if n < r{
				fmt.Println("입력하신 숫자가 더 작습니다..")
			} else{
				fmt.Println("숫자를 맞췄습니다. 시도한 횟수:",cnt)
				break
			} 
			cnt ++
		}
	}
}