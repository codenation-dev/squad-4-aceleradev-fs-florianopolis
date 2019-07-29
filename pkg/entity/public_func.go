package entity

// PublicFunc structures the data
type PublicFunc struct {
	CompleteName string  `json:"complete_name"`
	ShortName    string  `json:"short_name"`
	Wage         float64 `json:"wage"`
	Departament  string  `json:"departament"`
	Function     string  `json:"function"`
	Relevancia   int     `json:"relevancia"`
}

type PublicStats struct {
	Floor       string  `json:"floor"`
	Avg         float64 `json:"avg"`
	Qtd         int     `json:"qtd"`
	Departament string  `json:"departament"`
	Cargo       string  `json:"cargo"`
	Min         float64 `json:"min"`
	Max         float64 `json:"max"`
}
