package pac3

import (
	"fmt"
	"math"
)

type Sphere struct {
	Radius float64
}

func (s Sphere) Volume() float64 {
	return (4 * math.Pi * math.Pow(s.Radius, float64(3))) / 3
}

func (s Sphere)String()string{
	return fmt.Sprintf("It is a Sphere of radius %v \n", s.Radius)
}