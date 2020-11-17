package chapter13

import (
	"fmt"
	"time"
)

func a() {
	for i := 0; i < 50; i++ {
		fmt.Print("a ")
	}
}

func b() {
	for i := 0; i < 50; i++ {
		fmt.Print("b ")
	}
}

func Test01() {
	fmt.Println("Start Test01()")
	go a()
	go b()
	time.Sleep(time.Second)
	fmt.Println("end Test01()")
}
