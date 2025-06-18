package model

type Response struct {
	ResponseCode string            `json:"responseCode"`
	Message      string            `json:"message"`
	Errors       map[string]string `json:"errors,omitempty"`
	Data         any               `json:"data,omitempty"`
	Token        string            `json:"token,omitempty"`
}
