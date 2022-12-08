package main

import "fmt"

func rangee() {
	var arr = [3]int{1, 2, 3}
	var index int
	var value int
	for index, value = range arr {
		fmt.Printf("%v %v \n", index, value)
	}
}
