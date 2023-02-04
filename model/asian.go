package model

import "fmt"

type Asian struct {
	Height int
	Name   string
}

func (w Asian) speak() {
	fmt.Println("Im Asian", w.Name)
}
