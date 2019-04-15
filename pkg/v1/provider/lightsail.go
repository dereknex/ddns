package provider

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lightsail"
	"github.com/aws/aws-sdk-go/service/lightsail/lightsailiface"
)

// LightsailProvider is a DNS provider
type LightsailProvider struct {
	svc lightsailiface.LightsailAPI
}

// InitLightsailProvider create a lightsail struct
func InitLightsailProvider(region string) (*LightsailProvider, error) {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(region)})
	if err != nil {
		return nil, err
	}
	svc := lightsail.New(sess)
	return &LightsailProvider{
		svc: svc,
	}, nil
}

// GetRecord read a dns record
func (l *LightsailProvider) GetRecord(record string) (value string, err error) {

	return "", nil
}
