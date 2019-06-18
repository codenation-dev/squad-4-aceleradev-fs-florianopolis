package models

// Customer of the bank
type Customer struct {
	ID          int     `json:"cid"`
	Name        string  `json:"name"`
	Wage        float32 `json:"wage"`
	IsPublic    int8    `json:"is_public"`
	SentWarning string  `json:"sent_warning"` //TODO: Isso pode se tornar o ID da tabela warning
}
