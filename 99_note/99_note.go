// this file just for note, not for run
// 這個檔案只是用來做筆記的，不是用來執行的，所以裡面可能會有一些錯誤的程式碼，請不要直接執行這個檔案，

// go run .\00_basic_study.go
/* go build .\00_basic_study.go >>>> .exe file */

// a-z, A-Z, 0-9, and _
// const 全大階
// type, func, var Snake Case 全小階_

package main

import (
	"fmt"
	"math"
)

// type 
int      	// >>>> 0 
int, int8, int16, int32 (rune), int64      // int default 係32 bits, 64 bits based on system
uint, uint8 (byte), uint16, uint32, uint64 // uint default 係32 bits, 64 bits based on system
float32		// >>>> 0
float64		// >>>> 0
map[string]int // >>>> map[]


const (
	JAVA = iota // 0
	PYTHON      // 1
)


// %v Prints the value in a default format                >>>> Hello World
// %#v Prints the value in Go-syntax format `Hello World` >>>> "Hello \n World"
// %T Prints the type of the value                        >>>> string
Printf("i has value: %v and type: %T\n", i, i) 


// -------------------------------------------------------------


// 要pointer receiver 才能修改 struct 的值，否則只能修改副本
// 不同java, 沒有class, 不是class裡的method
// 但 employee又不用 pass address &
func (e *Employee) Set_Skills(s Skills) {
    e.Skills = s
}

func main() {
	Bob := Employee{Person{"Bob", 30}, JAVA, 30}
    // 不用&
	Bob.Set_Skills(PYTHON)
	// 直接賦值唔安全
    Bob.Skills = PYTHON
}