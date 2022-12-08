package main

import "fmt"

func forr() [5]int {
	fmt.Println("Implemetation for ")
	var arr = [5]int{1, 2, 3, 4 ,5}
	var i int
	for i = 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	return arr

}
