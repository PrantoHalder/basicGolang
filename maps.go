package main

import "fmt"

func enamul() {
	var a = map[string]string{"name": "pranto", "Job": "Lecturer"}
	fmt.Println(a)

	var b = make(map[string]string)
	b["name"] = "pranto"
	fmt.Println(b["name"])
	b["name"] = "Shovon"
	fmt.Println(b["name"])
	delete(b, "name")
	fmt.Println(b)

	var val1 string
	var val2 string
	var ok1 bool
	var ok2 bool

	val1, ok1 = a["name"]
	val2, ok2 = a["Job"]
	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
    var k string
	var v string
	for k,v = range a {
		fmt.Println(k,v)
	}

}
