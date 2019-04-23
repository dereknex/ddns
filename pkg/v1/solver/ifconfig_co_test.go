package solver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfconfigCoRun(t *testing.T) {
	s := &IfconfigCo{}
	ip, err := s.Run()
	assert.Nil(t, err)
	assert.NotEqual(t, ip, "")
	assert.True(t, len(ip) < 16, ip)
}
