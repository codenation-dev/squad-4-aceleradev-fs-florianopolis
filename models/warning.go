package models

// Warning Models the warning
type Warning struct {
	WID          int32  `json:"WID"`
	Dt           string `json:"dt"`
	Message      string `json:"message"`
	SentTo       string `json:"sent_to"`
	FromCustomer string `json:"from_customer"`
}
