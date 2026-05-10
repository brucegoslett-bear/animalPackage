package main

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	name, age = "Kim", 23

	a1 = iota
	a2
	a3
	a4
)

func init() {
	fmt.Printf("\n\n\n")
	fmt.Println("The land before time!")
	fmt.Printf("3\t2\t1\t\n\n\n")
}

func main() {

	fmt.Println("Hello world")

	x := rand.Intn(101)
	fmt.Printf("%v is %v years old today and her number is %v \n\n", name, age, x)
	var y float64 = float64(x)
	var z float64 = math.Sqrt(y)
	fmt.Println("Her problem number is", z)
	fmt.Println()

	var h, i, j, nuTrip, nuDoub int
	var repetitions int = 9
	var loopCount int = 0
	var sum, greatSum, greatLoop int
	for c := 0; c <= repetitions; c++ {
		loopCount++
		fmt.Println("This is Roll", loopCount)
		for b := 1; b <= 3; b++ {
			a := rand.Intn(6) + 1
			if b == 1 {
				h = a
				fmt.Printf("On dice %v you rolled a %v.\n", b, a)
			}
			if b == 2 {
				i = a
				fmt.Printf("On dice %v you rolled a %v.\n", b, a)
			}
			if b == 3 {
				j = a
				fmt.Printf("On dice %v you rolled a %v. \n Well done!!\n", b, a)
				sum = (h + i + j)
				if sum > greatSum {
					greatSum = sum
					greatLoop = loopCount
				}
				fmt.Printf("You rolled %d+%d+%d Your Total throw is %v\n\n", h, i, j, sum)
				if h == i && h == j {
					fmt.Printf("\n\n\t\t\t\t\t\t\t-!!!Tripple!!!-\n\n")
					nuTrip++
				} else if h == i || h == j || i == j {
					fmt.Printf("\n\n\t\t~!!DOUBLE!!~\n\n")
					nuDoub++
				}
			}

		}
		if c == repetitions {
			fmt.Printf("\tYou rolled %v Tripps and %v Doubles\n\n", nuTrip, nuDoub)
			fmt.Printf("\t\t✨✨✨The legendary Roll was in loop %v with a grand total of %v✨✨✨\n\n", greatLoop, greatSum)
		}
	}

	fmt.Println()
	for e := 0; e < 10; e++ {
		d := rand.Intn(150) + 1
		fmt.Println(d)
		switch {
		case d == 100:
			fmt.Println("d is 100")
		case d <= 50:
			fmt.Println("d is less than 50")
		default:
			fmt.Println("d is greater than 50")
		}
	}
	hexloopcount := 0
	fmt.Println("These are the things...")
	for hexloopcount = 0; hexloopcount < repetitions; hexloopcount++ {
		fmt.Printf(" %T\t%v\t%b\t%#x\t\n", a1, a1, a1, a1)
		fmt.Printf(" %T\t%v\t%b\t%#x\t\n", a2, a2, a2, a2)
		fmt.Printf(" %T\t%v\t%b\t%#x\t\n", a3, a3, a3, a3)
		fmt.Printf(" %T\t%v\t%b\t%#x\t\n", a4, a4, a4, a4)
	}
}
