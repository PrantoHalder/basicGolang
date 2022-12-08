package main

import "fmt"

func nested() {

	fmt.Println("implemantation of nested for loop")
	var arr = [3]string{"pranto", "halder", "shovon"}
	var arr1 = [3]string{"apple", "orange", "malta"}
	var i int
	var j int
	for i = 0; i < len(arr); i++ {
		for j = 0; j < len(arr1); j++ {
			fmt.Println(arr[i], arr1[j])
		}

	}
}
