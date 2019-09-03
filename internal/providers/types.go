package providers

import "google.golang.org/api/gmail/v1"

type Provider interface {
	Login()
	GetEmails(pageSize int64, pageId string) (GetEmailResponse, error)
}

type GmailProvider struct {
	Service *gmail.Service
	Provider
}

type GetEmailResponse struct {
	ResultSize 	int64	  	`json:"resultSize"`
	NextPage 	string 		`json:"nextPage"`
	PrevPage	string  	`json:"prevPage"`
	Results 	string	`json:"results"`
}
