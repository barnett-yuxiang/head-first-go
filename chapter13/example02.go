package chapter13

import "fmt"

func Test02() {
	var myChannel1 chan float64
	myChannel1 = make(chan float64)
	fmt.Println("xx", myChannel1)
	//myChannel1 <- 3.14

	myChannel2 := make(chan string)
	go greeting(myChannel2)
	fmt.Println(<-myChannel2)

}

func greeting(myChannel chan string) {
	myChannel <- "hi"
}
