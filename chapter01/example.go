package chapter01

import (
	"fmt"
	"reflect"
)

func Test01() {
	var length float64 = 1.2
	var width int = 2
	fmt.Println(length * float64(width))

	fmt.Println(reflect.TypeOf(length))
	fmt.Println(reflect.TypeOf(width))
	fmt.Println(reflect.TypeOf(float64(width)))
}
