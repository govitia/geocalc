package geocalc

import "math"

// All these formulas are for calculations on the basis of a spherical earth
// (ignoring ellipsoidal effects) – which is accurate enough for most purposes…
// [In fact, the earth is very slightly ellipsoidal;
// using a spherical model gives errors typically up to 0.3%].
// https://www.movable-type.co.uk/scripts/latlong.html

// EarthRadius in meters
const EarthRadius = 6371e3

// Distance return the distance between two earth Point in meters
func Distance(p1, p2 Point) float64 {
	phi1 := p1.Lat * math.Pi / 180 // phi, alpha in radians
	phi2 := p2.Lat * math.Pi / 180
	deltaPhi := (p2.Lat - p1.Lat) * math.Pi / 180
	deltaAlpha := (p2.Lon - p1.Lon) * math.Pi / 180

	a := math.Sin(deltaPhi/2)*math.Sin(deltaPhi/2) +
		math.Cos(phi1)*math.Cos(phi2)*
			math.Sin(deltaAlpha/2)*math.Sin(deltaAlpha/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return EarthRadius * c // in metres
}
