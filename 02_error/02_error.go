package main

import (
	"fmt"
)

func throwsPanic(f func()) (b string) {
	// recover() 只能在 defer 裡生效
	// recover() 類似java的try-catch.

	defer func() {
		if x := recover(); x != nil {
			b = "hello, recover"
		}
	}()
	f()    //執行函式 f，如果 f 中出現了 panic，那麼就可以恢復回來
	return //隱式return b, b=="" || b=="hello, recover"
}

func main() {
	ok := throwsPanic(func() {
		// panic 用於引發一個恐慌（panic）狀態，這通常表示程式遇到了嚴重的錯誤或異常情況，無法繼續正常執行。
		panic("something went wrong")
	})
	fmt.Println(ok) // >>>> hello, recover
}
