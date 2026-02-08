// this file just for note, not for run
// 這個檔案只是用來做筆記的，不是用來執行的，所以裡面可能會有一些錯誤的程式碼，請不要直接執行這個檔案，

// go run .\00_basic_study.go
/* go build .\00_basic_study.go >>>> .exe file */

// a-z, A-Z, 0-9, and _
// const 全大階
// var Snake Case 全小階_

package main

// import 寫法
import (
	"fmt"
	"math"
)

// type 
int      	// >>>> 0 
int, int8, int16, int32, int64      // int default 係32 bits, 64 bits based on system
uint, uint8, uint16, uint32, uint64 // uint default 係32 bits, 64 bits based on system
rune 		// >>>> rune 是 int32 的別名
byte 		// >>>> byte 是 uint8 的別名

float32		// >>>> 0
float64		// >>>> 0
string 		// >>>> ""
bool		// >>>> false
map[string]int // >>>> map[]

const (
	JAVA = iota // 0
	PYTHON      // 1
)
type Skills byte
type Person struct {
	name string
	age  int
}
type Employee struct {
    Person
	Skills
    age int
}

func (e Employee) getInfo() (name string, age int) {
	return e.name, e.age
}

func (s Skills) String() string {
    strings := []string {"JAVA", "PYTHON"}
    return strings[s]
}

// 要pointer receiver 才能修改 struct 的值，否則只能修改副本
// 不同java, 沒有class, 不是class裡的method
// 但 employee又不用 pass address &
func (e *Employee) Set_Skills(s Skills) {
    e.Skills = s
}

func main() {
    // Bob := Employee{Person:Person{name:"Bob", age:30}, Skills:Skills{"Java"}, age:30}
    // fmt.Print(Bob.age, Bob.Person.age) // >>>> 30 30

	Bob := Employee{Person{"Bob", 30}, JAVA, 30}
	fmt.Println(Bob.getInfo()) // >>>> Bob30
	fmt.Println(Bob.Skills.String()) // >>>> JAVA

	// 
    Bob.Set_Skills(PYTHON)
    fmt.Println(Bob.Skills.String()) // >>>> PYTHON

	// 直接賦值唔安全
    Bob.Skills = PYTHON
    fmt.Println(Bob.Skills.String()) // >>>> PYTHON

}




type type_func func(int) bool
func x_func(y ...int) (int) {return y[0]*2}
const x int = 1
var x_var int = 1
var x_arr [5]int = [5]int{1, 2, 3, 4, 5}


func forloop(array ...int){ // ...int == 
	for _index, value := range array {
		fmt.Println(_index, value)
	}
}



func throwsPanic(f func()) (b bool) {
	// recover() 只能在 defer 裡生效
	// recover() 類似java的try-catch. 

    defer func() {
        if x := recover(); x != nil { 
            b = true
        }
    }()
    f() //執行函式 f，如果 f 中出現了 panic，那麼就可以恢復回來
    return //隱式return b, b==false
} 

// panic 是 Go 語言中的一個內建函式，用於引發一個恐慌（panic）狀態，這通常表示程式遇到了嚴重的錯誤或異常情況，無法繼續正常執行。
func init() {
    if user == "" {
        panic("no value for $USER") 
    }
}

func main(){
	// %v, %#v, %T, %%
	// %v Prints the value in a default format                >>>> Hello World
	// %#v Prints the value in Go-syntax format `Hello World` >>>> "Hello \n World"
	// %T Prints the type of the value                        >>>> string
	// %%                                                     >>>> %
	Printf("i has value: %v and type: %T\n", i, i) 

	make([]int, 5, 10) // make([]type, length, capacity)
	myslice1 = append(myslice1, 20, 21, 22, 23) // Change length by appending items
	myslice3 := append(myslice1, myslice2...)
	copy(numbersCopy, neededNumbers)

}