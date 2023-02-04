package product

import "udemy/go/factory"

type Vittsjo struct {
}

func (l *Vittsjo) Price() float64 {
	return 899000
}

func (v *Vittsjo) Size() factory.Dimension {
	return factory.Dimension{
		80,
		70,
		40,
	}
}
func (v *Vittsjo) IsFoldable() bool {
	return false
}
