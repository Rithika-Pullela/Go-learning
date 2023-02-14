package pac1

import (
	"fmt"
)

type Cube struct {
	Length float64
}

func (c Cube) Volume() float64 {
	return c.Length * c.Length * c.Length
}

func (c Cube)String()string{
	return fmt.Sprintf("It is a cube of length %v \n", c.Length)
}
