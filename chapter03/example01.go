package chapter03

import "fmt"

func Test01() {
	var width, height, area float64
	width = 4.2
	height = 3.0
	area = width * height
	fmt.Printf("%0.2f liters needed\n", area/10.0)
	width = 5.2
	height = 3.5
	area = width * height
	fmt.Println(area/10.0, "liters needed")
	sayHi()
	fmt.Printf("double 2.0 = %.2f\n", double(2.0))
}

func sayHi() {
	fmt.Println("Hi")
}

func double(number float64) float64 {
	return number * 2
}
