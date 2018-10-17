package examples

import (
	"reflect"
	"testing"
)

func TestCountCharacterNum(t *testing.T) {

	r := CountCharacterNum("Hello World!")
	if reflect.DeepEqual(r, map[string]int{"": 1, "H": 1, "e": 1, "o": 2, "W": 1, "!": 1, "r": 1, "l": 3}) {
		t.Log("Success")
	}
}
