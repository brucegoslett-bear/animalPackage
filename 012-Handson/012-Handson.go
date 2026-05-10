package main

import (
	"fmt"

)

func main() {

	x := 43
	y := 40
	z := 45

	fmt.Printf(" x=%v \t %T\t\n", x, x)

	if x > y && x < z {
		fmt.Printf("between %v\t %v\t\n", y, z)
	}

	if x > 42 || x > 43 {
		fmt.Println("greater")
	} else {
		fmt.Println("less")
	}

	if x != 42 {
		fmt.Println("not the end of the galaxy")
	}

	switch {
	case x == 42:
		fmt.Println("life oh life")
	case x > 42:
		fmt.Println("we went too far")
	default:
		fmt.Println("it was not a cow")
	}
}
