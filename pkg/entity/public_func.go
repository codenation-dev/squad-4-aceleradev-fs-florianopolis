package entity

// PublicFunc structures the data
type PublicFunc struct {
	CompleteName string  `json:"complete_name"`
	ShortName    string  `json:"short_name"`
	Wage         float64 `json:"wage"`
	Departament  string  `json:"departament"`
	Function     string  `json:"function"`
}
