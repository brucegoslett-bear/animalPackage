package main

import "fmt"

var a int = 1

const name, age = "Kim", 23

func main() {

	fmt.Println("her name is", name, "and her age is", age)
	fmt.Printf("the number %v is an %T\n", a, a)
	b := 23.0
	fmt.Printf("the number %v is an %T\n", b, b)
	fmt.Println(b + 19)
}
