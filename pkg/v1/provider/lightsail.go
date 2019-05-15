package provider

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lightsail"
	"github.com/aws/aws-sdk-go/service/lightsail/lightsailiface"
)

// LightsailProvider is a DNS provider
type LightsailProvider struct {
	svc     lightsailiface.LightsailAPI
	domains map[string]*lightsail.Domain
}

// InitLightsailProvider create a lightsail struct
func InitLightsailProvider(region string) (*LightsailProvider, error) {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(region)})
	if err != nil {
		return nil, err
	}
	svc := lightsail.New(sess)
	return &LightsailProvider{
		svc:     svc,
		domains: make(map[string]*lightsail.Domain),
	}, nil
}

// GetRecord read a dns record
func (l *LightsailProvider) GetRecord(record string, domain string) (value string, err error) {
	if l.domains[domain] != nil {
		for _, entry := range l.domains[domain].DomainEntries {
			if record == *entry.Name {
				value = *entry.Name
				return
			}
		}
	}
	input := &lightsail.GetDomainInput{
		DomainName: &domain,
	}
	output, err := l.svc.GetDomain(input)
	if err != nil {
		return
	}

	l.domains[domain] = output.Domain
	for _, entry := range l.domains[domain].DomainEntries {
		if record == *entry.Name {
			value = *entry.Target
			return
		}
	}
	return "", errors.New("not found")
}

// SetRecord change dns record, if record not exists create it
func (l *LightsailProvider) SetRecord(name, value string) (err error) {
	return err
}
