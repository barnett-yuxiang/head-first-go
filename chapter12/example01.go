package chapter12

import "fmt"

func Test01() {
	Socialize()
}

func Socialize() {
	defer fmt.Println("Good Bye")
	fmt.Println("Hello")
	fmt.Println("Nice Weather, eh?")
}

func Test02() {
	arrays := [2]string{"h", "e"}
	for index, item := range arrays {
		fmt.Println(index+1, item)
	}

	//for i := 1; i < 3; i++ {
	//	fmt.Println(arrays[i])
	//}
}

func forcePanic() {
	panic("oh, no, we're going down")
}

func calmDown() {
	//fmt.Println(recover())
	recover()
}

func Test03() {
	defer calmDown()
	forcePanic()
	fmt.Println("End of test03")
}
