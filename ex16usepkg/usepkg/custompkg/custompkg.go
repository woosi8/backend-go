package custompkg

import "fmt"

type Student struct{ //타입도 마찬가지로 대문자로 시작하면 외부 공개
	Name string
	Are int
	score int
}



func PrintCustom()  {
	fmt.Println("This is cusom package!")
	
}
// 함수 시작을 소문자로 하면 비공개 (외부에서 쓸수 없다)
func printcustom2()  {
	fmt.Println("This is cusom package2222!")

}