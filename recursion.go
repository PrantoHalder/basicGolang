package main

import "fmt"

func recursion(x int) int {
	if x==11{
		fmt.Println("it stop,s")
          return 0
	}
	fmt.Println(x)
	return recursion(x+1)


}
func pranto (){
	
	fmt.Println(recursion(2))

}