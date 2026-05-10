package main

import (
	"fmt"
	"math/rand"
)

func main() {

	x := rand.Intn(10)
	y := rand.Intn(10)

	/*	fmt.Println("this is x y", x, y)
		if x < 4 && y < 4 {
			fmt.Println("x and y are both less than 4")
		} else if x > 6 && y > 6 {
			fmt.Println("x and y are both greater than 6")
		} else if x >= 4 && x <= 6 {
			fmt.Println("x is greater than or equal to 4 and less than or equal to 6")
		} else if x != 6 {
			fmt.Println("x is not equal to 6")
		} else {
			fmt.Println("not true of any conditions")
		}

	*/
	fmt.Println("this is x y", x, y)
	switch {
	case x < 4 && y < 4:
		fmt.Println("x and y are both less than 4")
	case x > 6 && y > 6:
		fmt.Println("x and y are both greater than 6")
	case x >= 4 && x <= 6:
		fmt.Println("x is greater than or equal to 4 and less than or equal to 6")
	case y != 6:
		fmt.Println("y is not equal to 6")
	default:
		fmt.Println("not true of any conditions")

	}
}
