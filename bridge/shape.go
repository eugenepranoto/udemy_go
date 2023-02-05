// menjembatani 2 object missal ada kotak. lalu kita mau buat
// kotak biru, kotak merah, dll
// solusi tambah 1 props di kotak yang isinya object color

package bridge

import "fmt"

type Shape interface {
	Draw()
	Area() float64
}

type Circle struct {
	Radius int
	Color  Color
}

func (c Circle) Draw() {
	fmt.Println("drawing circle")
}

func (c Circle) Area() float64 {
	return 3.14 * float64(c.Radius) * float64(c.Radius)
}

type Square struct {
	Length int
}

func (s Square) Draw() {
	fmt.Println("drawing square")
}

func (s Square) Area() float64 {
	return float64(s.Length) * float64(s.Length)
}

func init() {
	c := Circle{}
	c.Draw()

	redCircle := Circle{Color: Red{}}
	redCircle.Draw()
}
