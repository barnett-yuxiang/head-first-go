package chapter04

import (
	"fmt"
	"head-first-go/chapter04/keyboard"
)

func Test02() {
	value, err := keyboard.GetFloat()
	fmt.Println(value, err)
}
