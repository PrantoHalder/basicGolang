package main

import "fmt"
type Persons struct{
	Name string
	age int

}
func Person (per Persons){
	fmt.Println("Name",per.Name)
	fmt.Println("age",per.age)
}
func St_function () {
    var per1 Persons
	var per2 Persons

	per1.Name ="pranto"
	per1.age=25

	per2.Name = "shovon"
	per2.age = 25

	Person(per1)
	Person(per2)

}