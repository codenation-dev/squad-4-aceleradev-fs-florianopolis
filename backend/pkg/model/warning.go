// Package model implements the models to structure the data
package model

// Warning models the warnings sent to users about the customers
type Warning struct {
	ID           int    `json:"id"`
	Dt           string `json:"dt"` //TODO: usar datetime
	Message      string `json:"msg"`
	SentTo       string `json:"sent_to"`       //TODO: usar id do user
	FromCustomer string `json:"from_customer"` //TODO: usar id do customer
}
