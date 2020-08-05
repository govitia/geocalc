package geocalc_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	geo "gitlab.com/psns/geocalc"
)

func TestNewPoint(t *testing.T) {
	lat, lon := float64(45), float64(54)
	p := geo.NewPoint(lat, lon)
	assert.NotNil(t, p)
	assert.Equal(t, lat, p.Lat)
	assert.Equal(t, lon, p.Lon)
}

func TestPoint_Bearing(t *testing.T) {
	p1 := geo.NewPoint(50, 3)
	p2 := geo.NewPoint(51, 3)
	assert.NotNil(t, p1)
	assert.NotNil(t, p2)
	assert.Equal(t, float64(0), p1.Bearing(*p2))
	assert.Equal(t, math.Pi, p2.Bearing(*p1))
}