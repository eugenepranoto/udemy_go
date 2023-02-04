package product

import "udemy/go/factory"

type HemlingBy struct {
}

func (h *HemlingBy) Price() float64 {
	return 1798000
}
func (h *HemlingBy) Style() factory.SofaStyle {
	return factory.EuropeanStyle
}
