package chapter11

import "fmt"

type NoiseMaker interface {
	MakeSound()
}

type AnyThing interface {
}

type Whistle struct {
}

func (w Whistle) MakeSound() {
	fmt.Println("Whistle make sound")
}

func AcceptAnyThing(thing interface{}) {
	fmt.Println(thing)
	whistle, ok := thing.(Whistle)
	if ok {
		whistle.MakeSound()
	}
}

func Test01() {
	w := Whistle{}
	AcceptAnyThing(w)
}
