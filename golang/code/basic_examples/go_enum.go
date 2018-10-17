package basic_examples

import "fmt"

func main() {
	const (
		Sunday = iota
		Monday
		Tuesday
		Thursday
		Friday
		Saturday
		numberOfDays
	)

	fmt.Println(Saturday)
	fmt.Println(numberOfDays)

	const (
		bit0, mask0 = 1 << iota, 1<<iota - 1
		bit1, mask1
		// 사용하지 않는 열거형
		_, _
		bit3, mask3
	)

	fmt.Println(bit1)
	fmt.Println(mask1)

	fmt.Println(bit3)
	fmt.Println(mask3)

}
