package strategy

import "time"

type Road struct {
}

func (r Road) CalculateCost(trip TripDetail) float64 {
	return 100000 + trip.distance()*5000
}

const (
	defaultSpeed = 80 //kmh
)

func (r Road) CalculateETA(startAt time.Time, trip TripDetail) time.Time {
	t := trip.distance() / defaultSpeed
	var y int = int(t)
	duration := time.Duration(y) * time.Hour
	return startAt.Add(duration)
}
