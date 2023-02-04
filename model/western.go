package model

import "fmt"

type Western struct {
	Height int
	Name   string
}

func (w Western) speak() {
	fmt.Println("Im Western", w.Name)
}
