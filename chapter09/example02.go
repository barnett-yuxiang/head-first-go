package chapter09

import "fmt"

type MyType string

func (m MyType) sayHi() {
	fmt.Println("Hi")
}

func Test02() {
	value := MyType("Kamakura")
	value.sayHi()
}
