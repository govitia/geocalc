package geocalc_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	geo "gitlab.com/psns/geocalc"
)

func TestDistance(t *testing.T) {
	p1 := geo.NewPoint(50, 3)
	p2 := geo.NewPoint(50.09, 3)
	assert.NotNil(t, p1)
	assert.NotNil(t, p2)
	assert.InDelta(t, 10e3, geo.Distance(p1, p2), 10)
}
