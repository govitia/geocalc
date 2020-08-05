package geocalc

import "math"

// Point define an Earth point
type Point struct {
	Lat float64 // Latitude of the point
	Lon float64 // Longitude of the point
}

// NewPoint create and return an new Point from given coordinates
func NewPoint(lat, lon float64) *Point {
	return &Point{
		Lat: lat,
		Lon: lon,
	}
}

// Bearing return the bearing in rad between current and given Point
// Algorithm from https://www.movable-type.co.uk/scripts/latlong.html
func (p Point) Bearing(pt Point) float64 {
	y := math.Sin(pt.Lon - p.Lon) * math.Cos(pt.Lat)
	x := math.Cos(p.Lat)*math.Sin(pt.Lat) - math.Sin(p.Lat)*math.Cos(pt.Lat)*math.Cos(pt.Lon - p.Lon)
	return math.Atan2(y, x)
}
