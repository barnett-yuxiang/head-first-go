package chapter07

import (
	"example-head-first-go/chapter07/datafile"
	"fmt"
	"log"
)

func Test01() {
	lines, err := datafile.GetStrings("chapter07/votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
}

func Test02() {
	var ranks map[string]int
	ranks = make(map[string]int)
	fmt.Println(ranks)

	myMap := map[string]float64{"a": 1.2, "b": 5.6}
	fmt.Println(myMap)

	counters := map[string]int{"a": 3, "b": 1}
	var value int
	var ok bool
	value, ok = counters["a"]
	fmt.Println(value, ok)
	value, ok = counters["c"]
	fmt.Println(value, ok)
}
