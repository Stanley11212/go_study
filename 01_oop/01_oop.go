package main

import (
	"fmt"
)

// **concept of OOP in Go,
// 方便在容易組合成不同的interface
// 例如：所有人類都需要空氣，但有人食素，有人食肉
//      定義素食主義者的interface，肉食主義者的interface,
//      兩者都有呼吸function, 但素食主義者只有吃素的function，肉食主義者只有吃肉的function

// 只要實現了interface裡的method，便是這個interface的實例
// **問題是只看object難以識別屬於那一個interface

type Skills byte // byte = uint8

const (
	Java   = iota // 0
	Python        // 1
)

const (
	Exhale = iota // 0
	Inhale        // 1
)

// 由於interface沒有任何定義，
// 這時 universal_interface 可以是int=5, 又可以是string="hello", 可以是任何type的變量
var universal_interface interface{}

// 呼吸相關的 interface
type Human interface {
	Breathe()
	getUserInfo() (string, int)
}

// 素食主義者
type Vegetarian struct {
	name          string
	breatheStatus int
}

// 肉食主義者
type Carnivores struct {
	name          string
	breatheStatus int
}

// 要pointer receiver 才能修改 struct 的值，否則只能修改副本
// 不同java, 沒有class, 不是class裡的method
// 但 employee又不用 pass address &
func (v *Vegetarian) Breathe() {
	v.breatheStatus = (v.breatheStatus + 1) % 2 // 0->1, 1->0
}

func (v *Vegetarian) getUserInfo() (string, int) {
	return v.name, v.breatheStatus
}

func (c *Carnivores) Breathe() {
	c.breatheStatus = (c.breatheStatus + 1) % 2 // 0->1, 1->0
}

func (c *Carnivores) getUserInfo() (string, int) {
	return c.name, c.breatheStatus
}

func main() {
	vBob := Vegetarian{name: "vBob", breatheStatus: 0}
	cBob := Carnivores{name: "cBob", breatheStatus: 0}

	// vBob.breatheStatus = 1 // 不建議（破壞封裝）
	// fmt.Println("test1", vBob.breatheStatus)
	// vBob.breatheStatus = 0 // 不建議（破壞封裝）
	// fmt.Println("test2", vBob.breatheStatus)

	people := []Human{&vBob, &cBob}

	for range 5 {
		for _, person := range people { // vBob, cBob 輪流呼吸

			switch p := person.(type) { // if person is Human
			case Human:
				breathStr := map[int]string{0: "呼氣", 1: "吸氣"}
				p.Breathe()
				name, breatheStatus := p.getUserInfo()
				fmt.Printf("%s 正在：%s\n", name, breathStr[breatheStatus])
			default:
				fmt.Println("error")
			}
		}
	}
}
