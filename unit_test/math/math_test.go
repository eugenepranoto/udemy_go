package math_test

import (
	"testing"
	"time"
	. "udemy/go/unit_test/math"
)

// pakai test table untuk set of input and expected outcome
func TestAdd(t *testing.T) {
	testTable := []struct {
		a, b           int
		exectedOutcome int
	}{
		{a: 1, b: 2, exectedOutcome: 3},
		{a: 1, b: -2, exectedOutcome: -1},
		{a: -1, b: -2, exectedOutcome: -3},
		{a: 0, b: 0, exectedOutcome: 0},
	}

	for _, test := range testTable {
		result := Add(test.a, test.b)
		if result != test.exectedOutcome {
			t.Logf("Got %d but expect %d", result, test.exectedOutcome)
			t.Fail()
		}
	}
}

func TestAddWithHirarchical(t *testing.T) {
	a := 10
	t.Run("a=positif", func(t *testing.T) {
		t.Run("b=positif", func(t *testing.T) {
			result := Add(a, 5)
			if result != 15 {
				t.Logf("Got %d but expect %d", result, 15)
				t.Fail()
			}
		})
		t.Run("b=negatif", func(t *testing.T) {
			result := Add(a, -5)
			if result != 5 {
				t.Logf("Got %d but expect %d", result, 5)
				t.Fail()
			}
		})

	})
}

func TestNeedsToBeSkip(t *testing.T) {
	t.Skip("this test will be skip")
}

func TestCallToDb(t *testing.T) {
	if testing.Short() {
		t.Skip("skip because calling db is way to long")
	}
	<-time.After(3 * time.Second)
}
