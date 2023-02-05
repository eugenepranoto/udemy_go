package strategy

import (
	"fmt"
	"time"
)

type Interface interface {
	CalculateCost(trip TripDetail) float64
	CalculateETA(startAt time.Time, trip TripDetail) time.Time
}

type Location struct {
	CityName string
	Lat, Lng float64
}

type TripDetail struct {
	Origin, Destination Location
}

func (t TripDetail) distance() float64 {
	return 10
}

func init() {
	road := Road{}
	flight := Flight{wp: WPImpl{}}

	ctx := NewContext(road)
	t1 := ctx.CalculateETA(time.Now(), TripDetail{
		Origin: Location{
			CityName: "Bogor",
			Lat:      222.22,
			Lng:      222.22,
		},
		Destination: Location{
			CityName: "Cenkareng",
			Lat:      111.11,
			Lng:      111.11,
		},
	})
	fmt.Println(t1)

	ctx.SetStrategy(flight)

	t2 := ctx.CalculateETA(t1, TripDetail{
		Origin: Location{
			CityName: "Cengkarang",
			Lat:      111.11,
			Lng:      111.11,
		},
		Destination: Location{
			CityName: "Surabaya",
			Lat:      333.33,
			Lng:      333.33,
		},
	})

	fmt.Println(t2)
}
