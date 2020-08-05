package geocalc

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
