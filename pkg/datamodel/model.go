package datamodel

import "time"

type HealthResponse struct {
	TotalWebsites int           `json:"total_websites"`
	Success       int           `json:"success"`
	Failure       int           `json:"failure"`
	TotalTime     time.Duration `json:"total_time"`
}

type UrlStatus struct {
	url    string
	status bool
}

type VerifyAccessTokenReponse struct {
	ClientID         string `json:"client_id"`
	ExpiresIn        int    `json:"expires_in"`
	Scope            string `json:"scope"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}
