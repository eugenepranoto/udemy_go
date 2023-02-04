package ikea

import (
	"udemy/go/factory"
	"udemy/go/factory/ikea/product"
)

type Ikea struct {
}

func (i *Ikea) CreateChair() factory.Chair {
	return &product.Lefifarne{}
}

func (i *Ikea) CreateCoffeeTable() factory.CoffeeTable {
	return &product.Vittsjo{}
}

func (i *Ikea) CreateSofa() factory.Sofa {
	return &product.HemlingBy{}
}
