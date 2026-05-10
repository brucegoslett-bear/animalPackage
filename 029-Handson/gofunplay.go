package main

import "fmt"

const (
	_ = iota
	a
	b
	c
	d
	e
	f
)

func main() {

	fmt.Println("Hello you potato!!")
	fmt.Println()
	fmt.Printf("The variable generated in an Iota is %v\t and is Type %T\t with binary of %b\t and Hex of %#x\t\n", a, a, a, a)
	fmt.Printf("The variable generated in an Iota is %v\t and is Type %T\t with binary of %b\t and Hex of %#x\t\n", b, b, b, b)
	fmt.Printf("The variable generated in an Iota is %v\t and is Type %T\t with binary of %b\t and Hex of %#x\t\n", c, c, c, c)
	fmt.Printf("The variable generated in an Iota is %v\t and is Type %T\t with binary of %b\t and Hex of %#x\t\n", d, d, d, d)
	fmt.Printf("The variable generated in an Iota is %v\t and is Type %T\t with binary of %b\t and Hex of %#x\t\n", e, e, e, e)
	fmt.Printf("The variable generated in an Iota is %v\t and is Type %T\t with binary of %b\t and Hex of %#x\t\n", f, f, f, f)

}
