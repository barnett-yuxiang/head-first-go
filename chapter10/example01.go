package chapter10

import "fmt"

type Date struct {
	year int
}

func Test01() {
	date := Date{
		year: 2020,
	}

	fmt.Println(date)
}