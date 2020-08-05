package geocalc_test

import (
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