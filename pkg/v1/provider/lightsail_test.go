package provider

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitLightsailProvider(t *testing.T) {
	p, err := InitLightsailProvider("us-west-2")
	assert.Nil(t, err)
	assert.NotNil(t, p)
}

func TestGetRecord(t *testing.T) {
	p, err := InitLightsailProvider("us-west-2")
	assert.Nil(t, err)
	assert.NotNil(t, p)
	p.svc = new(LightsailAPI)
	value, err := p.GetRecord("test", "example.com")
	assert.Nil(t, err)
	assert.Equal(t, "10.0.0.1", value)

}

func TestLightsailProvider_SetRecord(t *testing.T) {
	p, err := InitLightsailProvider("us-west-2")
	assert.Nil(t, err)
	assert.NotNil(t, p)
	p.svc = new(LightsailAPI)
	err = p.SetRecord("test", "10.0.0.1", "example.com")
	assert.Nil(t, err)
}
