package builder

type House struct {
	numOfWindows int
	numOfDoors   int
}
type HouseBuilder struct {
	house House
}

func (b *HouseBuilder) BuildWindow() *HouseBuilder {
	if b.house.numOfWindows < 5 {
		b.house.numOfWindows++
	}
	return b
}

func (b *HouseBuilder) BuildDoor() *HouseBuilder {
	b.house.numOfDoors++
	return b
}
