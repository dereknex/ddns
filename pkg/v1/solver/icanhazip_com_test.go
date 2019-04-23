package solver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIcanhazipRun(t *testing.T) {
	s := &IcanhazipCom{}
	ip, err := s.Run()
	assert.Nil(t, err)
	assert.NotEqual(t, ip, "")
	assert.Truef(t, len(ip) < 16, "%s %d", ip, len(ip))
}
