package entity

// Customer of the bank
type Customer struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Wage        float32 `json:"wage"`
	IsPublic    int8    `json:"is_public"`
	SentWarning string  `json:"sent_warning"` //TODO: Isso pode se tornar o ID da tabela warning
}

// User of the app
type User struct {
	ID    int    `json:"id"`
	Login string `json:"login"`
	Email string `json:"email"`
	Pass  string `json:"pass"`
}

// Warning Models the warning
type Warning struct {
	ID           int    `json:"wid"`
	Dt           string `json:"dt"` //TODO: usar datetime
	Message      string `json:"msg"`
	SentTo       string `json:"sent_to"`       //TODO: usar id do user
	FromCustomer string `json:"from_customer"` //TODO: usar id do customer
}

// PublicFunc models the public employee profile
type PublicFunc struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Wage  float32 `json:"wage"`
	Place string  `json:"place"` // Place of work
}
