package code

import "fmt"

func Switch() {
	words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}

	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")

		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly the right length:", wordLen)

		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word!")
		}
	}
}

func Switch2() {
loop:
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			fmt.Println(i, "is even")
		case i%3 == 0:
			fmt.Println(i, "is divisible by 3 but no 2")
		case i%7 == 0:
			fmt.Println("exist the loop!")
			break loop
		default:
			fmt.Println(i, "is boring")
		}
	}
}
