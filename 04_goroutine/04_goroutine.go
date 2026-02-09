package main

import (
	"fmt"
)

func sum_channel(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

func for_test_range_channel(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
		c <- sum
	}
	close(c)

}

func main() {

	// vif := 3.14
	// cif := make(chan interface{})
	// cif <- vif
	// vif = (<-cif).(float64)

	arr := []int{7, 2, 8, -9, 4, 0}
	carr := make(chan int, 1)
	carr <- 1
	// carr <- 2
	// fmt.Println(<-carr, <-carr)
	fmt.Println("method0", <-carr)
	fmt.Println()

	go sum_channel(arr[:], carr)
	go sum_channel(arr[:len(arr)/2], carr)
	go sum_channel(arr[len(arr)/2:], carr)

	//method 1
	varr1, varr2, varr3 := <-carr, <-carr, <-carr
	fmt.Println("method1", varr1, varr2, varr3)
	fmt.Println()

	// method 2 range 有error
	carr = make(chan int, 100)
	go for_test_range_channel(arr, carr)
	for i := range carr {
		fmt.Println("method2", i)
	}

}
