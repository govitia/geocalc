package geocalc

import "math"

// Point define an Earth point
type Point struct {
	Lat float64 // Latitude of the point in rad
	Lon float64 // Longitude of the point in rad
}

// NewPoint create and return an new Point from given coordinates
// lat and lon in radians
func NewPoint(lat, lon float64) Point {
	return Point{
		Lat: lat,
		Lon: lon,
	}
}

// Bearing return the bearing in rad between current and given Point
func (p Point) Bearing(pt Point) float64 {
	y := math.Sin(pt.Lon-p.Lon) * math.Cos(pt.Lat)
	x := math.Cos(p.Lat)*math.Sin(pt.Lat) - math.Sin(p.Lat)*math.Cos(pt.Lat)*math.Cos(pt.Lon-p.Lon)
	return math.Atan2(y, x)
}

// Walk return destination Point given dist and bearing from start p Point.
// bearing in rad (clockwise from north) and dist in meters
func (p Point) Walk(bearing, dist float64) Point {
	lat := math.Asin( math.Sin(p.Lat)*math.Cos(dist/EarthRadius) +
		math.Cos(p.Lat)*math.Sin(dist/EarthRadius)*math.Cos(bearing))
	lon := p.Lon + math.Atan2(math.Sin(bearing)*math.Sin(dist/EarthRadius)*math.Cos(p.Lat), math.Cos(dist/EarthRadius) -
		math.Sin(p.Lat)*math.Sin(lat))
	return Point{Lat: lat, Lon: lon}
}