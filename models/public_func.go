package models

// PublicFunc Models the public employee profile
type PublicFunc struct {
	ID    int32   `json:"id"`
	Name  string  `json:"name"`
	Wage  float32 `json:"wage"`
	Local string  `json:"local"`
}
