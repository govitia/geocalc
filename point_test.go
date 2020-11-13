package geocalc_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	geo "github.com/govitia/geocalc"
)

func TestNewPoint(t *testing.T) {
	lat, lon := math.Pi, math.Pi
	p := geo.NewPoint(lat, lon)
	assert.Equal(t, lat, p.Lat)
	assert.Equal(t, lon, p.Lon)
}

func TestPoint_Bearing(t *testing.T) {
	p1 := geo.NewPoint(0, 0)
	p2 := geo.NewPoint(math.Pi, 0)
	assert.Equal(t, float64(0), p1.Bearing(p2))
	assert.Equal(t, math.Pi, p2.Bearing(p1))
}

func TestPoint_Walk(t *testing.T) {
	p1 := geo.NewPoint(0, 0)
	p2 := p1.Walk(math.Pi/2, geo.EarthCircumference/2)
	assert.InDelta(t, float64(0), p2.Lat, 1e-15)
	assert.Equal(t, math.Pi, p2.Lon)
	p3 := p2.Walk(0, geo.EarthCircumference/2)
	assert.InDelta(t, float64(0), p3.Lat, 1e-15)
	assert.Equal(t, 2*math.Pi, p3.Lon)
}

func TestPoint_Degree(t *testing.T) {
	p1 := geo.NewPoint(0, 0)
	lat, lon := p1.Degree()
	assert.Equal(t, float64(0), lat)
	assert.Equal(t, float64(0), lon)
	p1.Lat, p1.Lon = math.Pi, math.Pi/2
	lat, lon = p1.Degree()
	assert.Equal(t, float64(180), lat)
	assert.Equal(t, float64(90), lon)
}
