package main

import "fmt"
func addition(x string){
	fmt.Printf("%s am a good boy \n",x)
}

func add (p int ,q int) int{
	return p+q
}
func sub (e int ,f int)(result int){
	result = e-f
	return
}

func function() {
	fmt.Println("this is all about fuction")
	addition("I")
	fmt.Println(add(10,20))
	fmt.Println(sub(10,20))

}