package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	stdin := bufio.NewReader(os.Stdin)

	var a int
	var b int

	// Scanln의 출력값 n,err
	n, err := fmt.Scanln(&a, &b) //Scanln값을 읽어온다

	// err 변수가 nil이 아니면 즉, 올바르지 않는 메모리 주소(nil)
	// err가 nil이 아니면 (에러값이 있다는뜻)
	if err != nil{
		fmt.Println(err)
		stdin.ReadString('\n') //ReadString: os.Stdin(표준입력)에서 어떤 문자(\n 비우라는뜻)가 나올때까지 읽어라 라는 뜻
	} else{
		fmt.Println(n,a,b)
	}

	m, err := fmt.Scanln(&a, &b) 

	if err != nil{
		fmt.Println(err)
		stdin.ReadString('\n') 
	} else{
		fmt.Println(m,a,b)
	}
}