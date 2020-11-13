package geocalc

import "math"

// All these formulas are for calculations on the basis of a spherical earth
// (ignoring ellipsoidal effects) – which is accurate enough for most purposes…
// [In fact, the earth is very slightly ellipsoidal;
// using a spherical model gives errors typically up to 0.3%].
// https://www.movable-type.co.uk/scripts/latlong.html

const (
	// EarthRadius in meters
	EarthRadius = 6371e3
	// EarthCircumference in meters
	EarthCircumference = 2 * EarthRadius * math.Pi
)

// Distance return the distance between two earth Point in meters
func Distance(p1, p2 Point) float64 {
	deltaPhi := p2.Lat - p1.Lat
	deltaAlpha := p2.Lon - p1.Lon

	a := math.Sin(deltaPhi/2)*math.Sin(deltaPhi/2) +
		math.Cos(p1.Lat)*math.Cos(p2.Lat)*
			math.Sin(deltaAlpha/2)*math.Sin(deltaAlpha/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return EarthRadius * c // in metres
}

// Intermediate calculate an intermediate point at any fraction along the great circle path between
// p1 and p2. Start path at p1, if fraction=0, returned Point is p1. If fraction=1,
// returned Point is p2.
func Intermediate(p1, p2 Point, fraction float64) Point {
	return p1.Walk(p1.Bearing(p2), Distance(p1, p2)*fraction)
}

// DegreeToRad convert given arc in degree to rad
func DegreeToRad(deg float64) float64 {
	return deg * (math.Pi / 180)
}

// RadToDegree convert given arc in rad to degree
func RadToDegree(rad float64) float64 {
	return rad * (180 / math.Pi)
}
