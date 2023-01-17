package basic_examples

import (
	"fmt"
	"sort"
)

type Records []MyType

type MyType struct {
	Job  int `json:"Job"`
	Cost int `json:"Cost"`
}

func (r Records) Len() int           { return len(r) }
func (r Records) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r Records) Less(i, j int) bool { return r[i].Cost > r[j].Cost }

func sort1() {

	// custom slice 를 이용해 정렬하는 방법
	records := make(Records, 0, 1000)
	records = append(records, MyType{Job: 100, Cost: 234})
	records = append(records, MyType{Job: 101, Cost: 4000})
	records = append(records, MyType{Job: 102, Cost: 700})

	sort.Sort(records)
	fmt.Println(records)
}

func sort2() {

	// sort.Slice 를 이용해 정렬하는 방법 (단, stable 하지는 않는다.)
	// 참고 https://golang.org/pkg/sort/#Slice
	records := make([]MyType, 0, 1000)
	records = append(records, MyType{Job: 100, Cost: 234})
	records = append(records, MyType{Job: 101, Cost: 4000})
	records = append(records, MyType{Job: 102, Cost: 700})

	sort.Slice(records, func(i, j int) bool {
		return records[i].Cost > records[j].Cost
	})
	fmt.Println(records)
}
