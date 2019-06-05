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
func (l *LightsailProvider) GetRecord(name string, domain string) (value string, err error) {
	d, err := l.getDomain(domain)
	if err != nil {
		return
	}
	for _, entry := range d.DomainEntries {
		if name == *entry.Name {
			value = *entry.Target
			return
		}
	}
	return "", errors.New("not found")
}

// SetRecord change dns record, if record not exists create it
func (l *LightsailProvider) SetRecord(name, value string, domain string) (err error) {
	_, err = l.getDomain(domain)
	if err != nil {
		return err
	}

	err = l.createEntry(name, value, domain)
	return err
}

func (l *LightsailProvider) createEntry(name, value, domain string) (err error) {
	t := "A"
	input := &lightsail.CreateDomainEntryInput{
		DomainEntry: &lightsail.DomainEntry{
			Name: &name,
			Target: &value,
			Type: &t,
		},
		DomainName: &domain,
	}
	_, err = l.svc.CreateDomainEntry(input)
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
