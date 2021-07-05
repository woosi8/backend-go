package exinit

import "fmt"

var (
	//실행순서
	 a= c + b //1
	 b = f() //3 d=4에서 ++ 5가 된다
	 c = f() //2 f함수 실행 d++로 4가 된다
	 d = 3
)

// 초기화되면 init 함수가 실행된다 (init는 임포트 될때 한번만 실행되는 것이다. 다른곳에서 임폴트 된다고 또 실행되는건 아니다) 
// 맨처음 임폴트될때만 실행하면 끝난다
func init()  {
	d++ //4 d는 6이된다
	fmt.Println("exinit.init function, d")
}

func f() int  {
	d++
	fmt.Println("f() d:",d)
	return d
}

func PrintD(){
	fmt.Println("d:", d)
}
