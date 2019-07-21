package model

// PublicFunc models the public employee profile with relevant informations
type PublicFunc struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Wage  float32 `json:"wage"`
	Place string  `json:"place"` // Place of work
}
