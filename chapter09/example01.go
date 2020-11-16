package chapter09

type Liter float64
type Milliliter float64
type Gallons float64

func ToGallons(l Liter) Gallons {
	return Gallons(l * 0.264)
}

func Test01() {

}
