package advanced_examples

func CountCharacterNum(s1 string) map[string]int {

	var dict = make(map[string]int)

	for _, char := range s1 {
		dict[string(char)] += 1
	}

	return dict
}
