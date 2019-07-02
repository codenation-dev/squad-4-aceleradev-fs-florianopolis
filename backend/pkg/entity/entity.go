// Package entity implements the models to structure the data
package entity

// Customer of the bank
type Customer struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" description:"Nome completo do cliente"`
	Wage        float32 `json:"wage" description:"Salário bruto mensal, sem os extras ocasionais (férias...)"`
	IsPublic    int8    `json:"is_public" description:"1- é funcionário público, 0- não é funcionário público"`
	SentWarning string  `json:"sent_warning" description:"Avisos enviados aos users"`
	//TODO: Isso pode se tornar o ID da tabela warning
}

// User of the app
type User struct {
	ID    int    `json:"id"`
	Login string `json:"login"` // TODO: retirar login, já implementei tudo com o email
	Email string `json:"email"`
	Pass  string `json:"pass"`
	// TODO: implementar uma opção de quais avisos receber (ex: salarios acima de 100k, estado de SP...)
}

// Warning models the warnings sent to users about the customers
type Warning struct {
	ID           int    `json:"id"`
	Dt           string `json:"dt"` //TODO: usar datetime
	Message      string `json:"msg"`
	SentTo       string `json:"sent_to"`       //TODO: usar id do user
	FromCustomer string `json:"from_customer"` //TODO: usar id do customer
}

// PublicFunc models the public employee profile with relevant informations
type PublicFunc struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Wage  float32 `json:"wage"`
	Place string  `json:"place"` // Place of work
}
