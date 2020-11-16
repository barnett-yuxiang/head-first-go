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
	amount = 2
	double(&amount)
	fmt.Printf("double 2.0 = %.2f\n", amount)
}

func sayHi() {
	fmt.Println("Hi")
}

func double(number *float64) float64 {
	*number = (*number) * 2
	return *number
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
