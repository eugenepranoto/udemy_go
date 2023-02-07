package iterator

import "fmt"

type Iterator interface {
	Next() int
	HasNext() bool
}

type IteratorCollection interface {
	CreateIterator(direction bool) Iterator
}

type Array []int

func (a Array) CreateIterator(direction bool) Iterator {
	return &ForwardIterator{Arr: a}
}

type ForwardIterator struct {
	Arr       Array
	CurrIndex int
}

func (f *ForwardIterator) Next() int {
	val := f.Arr[f.CurrIndex]
	f.CurrIndex++
	return val
}
func (f ForwardIterator) HasNext() bool {
	length := len(f.Arr)
	return f.CurrIndex <= length-1
}

func init() {
	arr := Array{1, 2, 3, 4}
	it := arr.CreateIterator(true)
	for it.HasNext() {
		fmt.Println(it.Next())
	}
}
