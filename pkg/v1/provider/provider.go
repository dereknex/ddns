package provider

// Provider for DNS service
type Provider interface {
	GetRecord(string, string) (string, error)
	SetRecord(string, string, string) error
}
