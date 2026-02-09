package main

import (
	"fmt"
	"reflect"
)

type I struct { // 建議型別名稱大寫開頭（exported）
	S string
	B bool
	I int
	F float32
}

func main() {
	var x I = I{"hello", true, 42, 3.14}

	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	// 直接 .Field()，不用 .Elem()
	tag := t.Field(0).Tag    // struct 的第一個欄位 tag
	s := v.Field(0).String() // 第一個欄位（string）的值
	b := v.Field(1).Bool()   // 第一個欄位（string）的值

	fmt.Println("t:", t)
	fmt.Println("v:", v)
	fmt.Println("type:", v.Type())
	fmt.Println("tag :", tag)
	fmt.Println("kind :", v.Kind())
	fmt.Println("s:", s)
	fmt.Println("b:", b)
	fmt.Println()

	var x2 float64 = 3.4
	v2 := reflect.ValueOf(x2)
	fmt.Println("type:", v2.Type())
	fmt.Println("kind is float64:", v2.Kind())
	fmt.Println("kind is float64:", v2.Kind() == reflect.Float64)
	fmt.Println("value:", v2.Float())
}
