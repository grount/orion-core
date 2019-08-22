package gmail

import "google.golang.org/api/gmail/v1"

type Provider interface {
	Login()
}

type Email interface {
	GetList() map[string] interface{}
}

type ProviderGmail struct {
	Service *gmail.Service
}
