package company

import "fmt"

type Employee interface {
	GetSalary() int
	TotalDivisionSalary() int
}

type CEO struct {
	Subordinates []Employee
}

func (c CEO) GetSalary() int {
	return 10
}

func (c CEO) TotalDivisionSalary() int {
	var sum int
	for _, s := range c.Subordinates {
		sum = sum + s.TotalDivisionSalary()
	}
	return c.GetSalary() + sum
}

type VP struct {
	Subordinates []Employee
}

func (v VP) GetSalary() int {
	return 8
}

func (v VP) TotalDivisionSalary() int {
	var sum int
	for _, s := range v.Subordinates {
		sum = sum + s.TotalDivisionSalary()
	}
	return v.GetSalary() + sum
}

type Kroco struct {
}

func (k Kroco) GetSalary() int {
	return 6
}

func (k Kroco) TotalDivisionSalary() int {
	return k.GetSalary()
}

func init() {
	k1 := &Kroco{}
	k2 := &Kroco{}
	k3 := &Kroco{}

	vpEng1 := VP{
		[]Employee{k1, k2},
	}

	vpEng2 := VP{
		[]Employee{k3},
	}

	Ceo := CEO{
		[]Employee{vpEng1, vpEng2},
	}

	fmt.Println(Ceo.TotalDivisionSalary())
}
