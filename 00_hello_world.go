// go run .\00_basic_study.go
// go build .\00_basic_study.go >> .exe file

package main
import ("fmt")

var a int
var d float64
// c := 20 // 這行會導致編譯錯誤，因為全局變量不能使用簡短聲明方式

func main(){
	c := 20
	fmt.Println("Hello World")
	fmt.Println("Welcome to Go Programming")
	fmt.Println("Value of a:", a)
	fmt.Println("Value of c:", c)
	fmt.Println("Value of d:", d)
}