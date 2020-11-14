package chapter02

import (
	"fmt"
	"time"
)

func Test01() {
	var now time.Time = time.Now()
	var year int = now.Year()
	fmt.Println(now)
	fmt.Println(year)
}

