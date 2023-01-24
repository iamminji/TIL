package code

import (
	"fmt"
	"testing"
)

func Test_divAndRemainder(t *testing.T) {
	x, y, z := divAndRemainder(5, 2)
	fmt.Println(x, y, z)
}

func Test_divAndRemainder2(t *testing.T) {
	x, y, z := divAndRemainder2(5, 2)
	fmt.Println(x, y, z)
}

func Test_divAndRemainder3(t *testing.T) {
	x, y, z := divAndRemainder3(5, 2)
	fmt.Println(x, y, z)
}
