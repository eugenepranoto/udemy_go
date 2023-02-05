package strategy

import "time"

type WeatherPredictor interface {
	Predict(cityName string) bool
}

type WPImpl struct {
}

func (w WPImpl) Predict(cityName string) bool {
	return false
}

type Flight struct {
	wp WeatherPredictor
}

func (f Flight) CalculateCost(trip TripDetail) float64 {
	return (10000000 + trip.distance()*50000) / 200
}

const (
	defaultAirSpeed = 500 //kmh
)

func (f Flight) CalculateETA(startAt time.Time, trip TripDetail) time.Time {
	t := trip.distance() / defaultAirSpeed
	var y int = int(t)
	duration := time.Duration(y) * time.Hour

	eta := startAt.Add(duration)
	if f.wp.Predict(trip.Origin.CityName) {
		eta = eta.Add(1 * time.Hour)
	}
	return eta
}
