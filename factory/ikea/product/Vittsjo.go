package product

import "udemy/go/factory"

type Vittsjo struct {
}

func (l *Vittsjo) Price() float64 {
	return 899000
}

func (v *Vittsjo) Size() factory.Dimension {
	return factory.Dimension{
		Length: 80,
		Width:  70,
		Height: 40,
	}
}
func (v *Vittsjo) IsFoldable() bool {
	return false
}
