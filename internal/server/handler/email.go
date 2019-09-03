package handler

import (
	"github.com/the-orion-team/orion-core/internal/providers"
	"net/http"
)

func EmailsHandler(writer http.ResponseWriter, request *http.Request) {
	p := providers.GetProvider(providers.ProviderGmailType)

	if p == nil {
		respondJSON(writer, http.StatusInternalServerError, nil)
	}
	res, err := p.GetEmails() // TODO: investigate warning

	if err != nil {
		respondJSON(writer, http.StatusInternalServerError, err)
	}

	respondJSON(writer, http.StatusOK, res)
}
