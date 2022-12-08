package main

import "fmt"

func breake() {
	var arr = [6]string{"apple", "orange", "grapes", "mango", "goava", "multa"}
	var i int
	fmt.Println("#break#")
	for i = 0; i < len(arr); i++ {

		if i == 3 {
			break
		}
		fmt.Println(arr[i])

	}
	fmt.Println("#continue#")
	for i = 0; i < len(arr); i++ {

		if i == 3 {
			continue
		}
		fmt.Println(arr[i])
	}

}
