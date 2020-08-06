package geocalc

import "math"

// All these formulas are for calculations on the basis of a spherical earth
// (ignoring ellipsoidal effects) – which is accurate enough for most purposes…
// [In fact, the earth is very slightly ellipsoidal;
// using a spherical model gives errors typically up to 0.3%].
// https://www.movable-type.co.uk/scripts/latlong.html

// EarthRadius in meters
const EarthRadius = 6371e3
const EarthCircumference = 2*EarthRadius*math.Pi

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

// DegreeToRad convert given arc in degree to rad
func DegreeToRad(deg float64) float64 {
	return deg*(math.Pi/180)
}

// RadToDegree convert given arc in rad to degree
func RadToDegree(rad float64) float64 {
	return rad*(180/math.Pi)
}