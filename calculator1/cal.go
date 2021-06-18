package main //시작점을 포함하는 패키지 main, 프로그램의 시작 포인트라는 뜻, 고언어의 메인패키지는 딱 하나만 존재한다. 시작점이니깐
// main 함수를 가지고 있는게 패키지 main
import "fmt" //fmt 다른 패키지 가져오기

func main(){ // main함수가 프로그램 시작점이라는 뜻 (프로그램이 로딩되면 main함수부터 시작된다)
	var a int = 10 // int 타입(정수)
	// 선언 대입문 a:= 10 이처럼 간략하게 선언할수 있다 위와같음
	var msg string ="hello Variable"

	a = 20
	msg = "Good Morning"
	fmt.Println(msg,a)
}

// 변수의 4가지 속성
// 이름
// 값
// 주소
// 타입 : 타입이 정해져야 변수의 사이즈가 정해진다.



