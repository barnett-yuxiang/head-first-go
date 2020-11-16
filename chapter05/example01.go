package chapter05

import "fmt"

func Test01() {
	var primes [5]int = [5]int{1, 2, 3, 4, 5}
	primes[0] = 2
	fmt.Println(primes)

	var note [5]string = [5]string{"aa", "bb", "cc", "dd", "ee"}
	fmt.Println(note)
	fmt.Printf("%#v\n", note)
	fmt.Printf("%#v\n", primes)

	for i := 0; i < len(note); i++ {
		fmt.Println(note[i])
	}

	for index, value := range primes {
		fmt.Println(index, value)
	}
}
