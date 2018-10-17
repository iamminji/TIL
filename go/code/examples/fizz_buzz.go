package examples

import "fmt"

func fizzBuzz() {
	for i := 0; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}
}

func BeerBottle99() {
	for i := 99; i > 0; i-- {
		fmt.Print(i)
		var b = " bottles"
		if i <= 1 {
			b = " bottle"
		}
		fmt.Println(b, " of beer on the wall,", i, b, "of beer.")
		fmt.Println("Take one down, pass it around, No more bottles of beer on the wall")
	}

	fmt.Println("No more bottles of beer on the wall, No more bottles of beer.")
	fmt.Println("Go to the store and buy some more, 99 bottles of beer on the wall.")
}

func main() {
	//fizzBuzz()
	BeerBottle99()
}
