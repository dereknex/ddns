package provider

import (
	"errors"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lightsail"
	"github.com/aws/aws-sdk-go/service/lightsail/lightsailiface"
)

// FetchInterval for lightsail server
const FetchInterval = time.Second * 180

type domainCache struct {
	domain  *lightsail.Domain
	fetchAt time.Time
}

// LightsailProvider is a DNS provider
type LightsailProvider struct {
	svc   lightsailiface.LightsailAPI
	items map[string]*domainCache
}

// InitLightsailProvider create a lightsail struct
func InitLightsailProvider(region string) (*LightsailProvider, error) {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(region)})
	if err != nil {
		return nil, err
	}
	svc := lightsail.New(sess)
	return &LightsailProvider{
		svc:   svc,
		items: make(map[string]*domainCache),
	}, nil
}

// GetRecord read a dns record
func (l *LightsailProvider) GetRecord(record string, domain string) (value string, err error) {

	for _, entry := range l.items[domain].domain.DomainEntries {
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

func (l *LightsailProvider) getDomain(domain string) (*lightsail.Domain, error) {
	item := l.items[domain]
	if item != nil {
		if item.fetchAt.Add(FetchInterval).Before(time.Now()) {
			return item.domain, nil
		}
	} else {
		item = &domainCache{fetchAt: time.Now()}
	}

	input := &lightsail.GetDomainInput{
		DomainName: &domain,
	}
	output, err := l.svc.GetDomain(input)
	if err != nil {
		return nil, err
	}
	item.domain = output.Domain
	l.items[domain] = item
	return output.Domain, nil
}
