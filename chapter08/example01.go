package chapter08

import "fmt"

func Test01() {
	var myStruct struct {
		number float64
		word   string
		toggle bool
	}

	myStruct.number = 3.14
	myStruct.word = "hello"
	myStruct.toggle = true

	fmt.Println(myStruct)
	fmt.Printf("%#v\n", myStruct)

	s := Subscriber{
		Name:   "yuli",
		Rate:   2.3,
		Active: false,
	}
	applyDiscount(&s)
	fmt.Println(s.Rate)
}

func applyDiscount(s *Subscriber) {
	s.Rate = 4.99
}

func Test02()  {
	employee := Employee{Name: "natalie"}
	employee.City = "GuangZhou"
	employee.Address.PostalCode = "020"
	fmt.Println(employee)
}