package provider

// Provider for DNS service
type Provider interface {
	GetRecord(string) (string, error)
}
