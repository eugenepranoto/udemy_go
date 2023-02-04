package bridge

type Color interface {
	Hex() string
}

type Red struct {
}

func (r Red) Hex() string {
	return "#83232"
}

type Green struct {
}

func (r Green) Hex() string {
	return "#392302"
}
