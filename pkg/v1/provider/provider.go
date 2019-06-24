package provider

import "errors"

// Provider for DNS service
type Provider interface {
	GetRecord(string, string) (string, error)
	SetRecord(string, string, string) error
}


func NewProvider(name string) (Provider, error) {
	switch name {
	case "lightsail":
		return InitLightsailProvider("us-east-1")
	}
	return nil, errors.New("invaild domain provider")
}