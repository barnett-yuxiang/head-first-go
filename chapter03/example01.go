package chapter03

import "fmt"

func Test01() {
	var width, height float64
	width = 4.2
	height = 3.0
	amount, err := paintNeeded(width, height)
	fmt.Println(err)
	fmt.Printf("%0.2f liters needed\n", amount)
	width = 5.2
	height = 3.5
	amount, err = paintNeeded(width, height)
	fmt.Println(err)
	fmt.Printf("%0.2f liters needed\n", amount)
	sayHi()
	fmt.Printf("double 2.0 = %.2f\n", double(2.0))
}

func sayHi() {
	fmt.Println("Hi")
}

func double(number float64) float64 {
	return number * 2
}

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", height)
	}
	area := width * height
	return area, nil
}
