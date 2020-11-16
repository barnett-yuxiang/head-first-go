package chapter06

import (
	"fmt"
	"os"
)

func Test01() {
	var notes []string
	notes = make([]string, 7)
	fmt.Println(notes)

	primes := make([]int, 5)
	primes[0] = 2
	primes[1] = 3
	fmt.Println(primes)

	array1 := [5]string{"a", "b", "c", "d", "e"}
	slice1 := array1[:3]
	array1[1] = "X"
	fmt.Println(slice1)
}

func Test02() {
	fmt.Println(os.Args[1:])
	myFunc(1, "2", "a")
}

func myFunc(param1 int, param2 ...string) {
	fmt.Println(param1)
	fmt.Println(param2)
}
