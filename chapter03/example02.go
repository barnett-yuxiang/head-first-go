package chapter03

import (
	"errors"
	"fmt"
)

func Test03() {
	err := errors.New("height can't be negative")
	fmt.Println(err)
}
