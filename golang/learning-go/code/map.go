package code

import "fmt"

func Map() {
	teams := map[string][]string{
		"Orcas":   []string{"Fread", "Ralph", "Bijou"},
		"Lions":   []string{"Sarah", "Peter", "Billie"},
		"Kittens": []string{"Waldo", "Raul", "Ze"},
	}

	fmt.Println(teams)

	ages := make(map[int][]string, 10)
	ages[3] = []string{"minji"}
	fmt.Println(ages)

	totalWins := map[string]int{}
	totalWins["Orcas"] = 1
	fmt.Println(totalWins)
	totalWins["Kittens"]++
	fmt.Println(totalWins)

	v, ok := totalWins["Lions"]
	fmt.Println(v, ok)

	delete(totalWins, "Orcas")
	fmt.Println(totalWins)
	delete(totalWins, "Lions")
	fmt.Println(totalWins)
}
