package provider

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

// Lightsail is a DNS provider
type Lightsail struct {
	sess *session.Session
}

// InitLightsail create a lightsail struct
func InitLightsail(region string) (*Lightsail, error) {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(region)})
	if err != nil {
		return nil, err
	}
	return &Lightsail{
		sess: sess,
	}, nil
}

// GetRecord read a dns record
func (l *Lightsail) GetRecord(record string) (value string, err error) {

	return "", nil
}
