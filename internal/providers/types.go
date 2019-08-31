package providers

import "google.golang.org/api/gmail/v1"

type Provider interface {
	Login()
	NewInstance() Provider
}

type Email interface {
	GetEmails() map[string] interface{}
}

type ProviderGmail struct {
	Service *gmail.Service
	Provider
}
