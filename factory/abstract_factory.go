package factory

type Pricey interface {
	Price() float64
}

type Chair interface {
	Pricey
	IsIotEnaled() bool
	IsSoft() bool
}

type Dimension struct {
	Length, Width, Height int
}

type CoffeeTable interface {
	Pricey
	Size() Dimension
	IsFoldable() bool
}

type SofaStyle string

const (
	EuropeanStyle SofaStyle = "european style"
	AsianStyle    SofaStyle = "asian style"
)

type Sofa interface {
	Pricey
	Style() SofaStyle
}

type FurnitureFactory interface {
	CreateChair() Chair
	CreateCoffeeTable() CoffeeTable
	CreateSofa() Sofa
}
