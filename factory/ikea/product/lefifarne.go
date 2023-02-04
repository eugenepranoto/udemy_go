package product

type Lefifarne struct {
}

func (l *Lefifarne) Price() float64 {
	return 1095000
}

func (l *Lefifarne) IsIotEnaled() bool {
	return false
}

func (l *Lefifarne) IsSoft() bool {
	return false
}
