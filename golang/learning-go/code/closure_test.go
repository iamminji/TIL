package code

import (
	"fmt"
	"testing"
)

func TestClosureExample(t *testing.T) {
	ClosureExample()
}

func TestMakeMult(t *testing.T) {
	twoBase := MakeMult(2)
	threeBase := MakeMult(3)
	for i := 0; i < 3; i++ {
		fmt.Println(twoBase(i), threeBase(i))
	}
}
