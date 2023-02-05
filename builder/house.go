package builder

// membuat struct untuk membuat suatu struct lain
// dikasus ini ada struct housebuild yang mereturn house

import (
	"fmt"
	"log"
)

type House struct {
	numOfWindows int
	numOfDoors   int
	HasGarage    bool
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

func (b *HouseBuilder) SetGarage(t string) *HouseBuilder {
	if t == "large" {
		b.house.HasGarage = true
	}
	return b
}

func (b *HouseBuilder) CreateHouse() (*House, error) {
	if !b.house.HasGarage {
		return nil, fmt.Errorf("No Garrage")
	}
	return &b.house, nil
}

func init() {
	hb := HouseBuilder{}
	house, err := hb.BuildWindow().BuildDoor().SetGarage("large").CreateHouse()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(house)
}
