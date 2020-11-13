package geocalc_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	geo "gitlab.com/psns/geocalc"
)

func TestDistance(t *testing.T) {
	p1 := geo.NewPoint(0, 0)
	p2 := geo.NewPoint(0, math.Pi)
	assert.NotNil(t, p1)
	assert.NotNil(t, p2)
	assert.Equal(t, geo.EarthCircumference/2, geo.Distance(p1, p2))
}

func TestDegreeToRad(t *testing.T) {
	assert.Equal(t, math.Pi, geo.DegreeToRad(180))
	assert.Equal(t, math.Pi/2, geo.DegreeToRad(90))
}

func TestRadToDegree(t *testing.T) {
	assert.Equal(t, float64(180), geo.RadToDegree(math.Pi))
	assert.Equal(t, float64(90), geo.RadToDegree(math.Pi/2))
}

func TestIntermediate(t *testing.T) {
	p1 := geo.NewPoint(0, 0)
	p2 := geo.NewPoint(0, math.Pi)
	assert.NotNil(t, p1)
	assert.NotNil(t, p2)
	p3 := geo.Intermediate(p1, p2, .5)
	assert.InDelta(t, float64(0), p3.Lat, 1e-15)
	assert.Equal(t, math.Pi/2, p3.Lon)
}
