package main

// C:\web\Go\ex16usepkg\usepkg>go mod tidy  하면 메인에있는 import안에 필요한것들을 다 다운받는다
import (
	"fmt"
	"goproject/usepkg/custompkg"
	"goproject/usepkg/exinit" //import하면 패키지가 초기화된다 처음에. 그안에 변수들이 초기화 된다.(exinit.go)

	"github.com/guptarohit/asciigraph"
	"github.com/tuckersGo/musthaveGo/ch16/expkg"
)

func main()  {
	custompkg.PrintCustom()
	exinit.PrintD()
	// custompkg.printcustom2()
	expkg.PrintSample()

	data := []float64{3,4,5,6,9,7,5,8,5,10,2,7,2,5,6}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)
}


