package main

import"fmt"

type PERSON struct {
	name string
	age int
	job string 
}
func halder(){

	var per1 PERSON
	var per2 PERSON
     
	per1.name = "Pranto"
	per1.age=30
	per1.job = "doctor"

	per2.name = "shovon"
	per2.age=30
	per2.job="Lecturer"

	fmt.Println("Name",per1.name)
	fmt.Println("age",per1.age)
	fmt.Println("Job",per1.job)

	fmt.Println("Name",per2.name)
	fmt.Println("age",per2.age)
	fmt.Println("Job",per2.job)


}