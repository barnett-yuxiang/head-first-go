package chapter03

import (
	"errors"
	"fmt"
	"reflect"
)

func Test03() {
	err := errors.New("height can't be negative")
	fmt.Println(err)
	myInt, myBool, myString := manyReturns()
	fmt.Println(myInt, myBool, myString)
}

func manyReturns() (int, bool, string) {
	return 1, true, "hello"
}

func Test04() {
	amount := 6
	fmt.Println(amount)
	fmt.Println(&amount)

	var pi float32
	fmt.Println(reflect.TypeOf(&amount))
	fmt.Println(reflect.TypeOf(&pi))
	fmt.Println(pi)

	var myBool bool
	myBoolPointer := myBool
	fmt.Println(myBoolPointer)
}
