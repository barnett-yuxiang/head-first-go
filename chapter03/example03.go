package chapter03

import "fmt"

func createPointer1(value float64) *float64 {
	tmp := &value
	fmt.Println(tmp)
	fmt.Println(&value)
	return tmp
}

func createPointer2(value float64) (*float64, *float64) {
	tmp := value + 1
	return &value, &tmp
}

func Test05() {
	var cp1 *float64 = createPointer1(123)
	fmt.Println(cp1)

	cp21, cp22 := createPointer2(123.0)
	fmt.Println(cp21)
	fmt.Println(cp22)
	fmt.Println(*cp22)
}
