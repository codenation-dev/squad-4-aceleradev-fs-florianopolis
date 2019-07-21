package adding

// Customer of the bank
type Customer struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" description:"Nome completo do cliente"`
	Wage        float32 `json:"wage" description:"Salário bruto mensal, sem os extras ocasionais (férias...)"`
	IsPublic    int8    `json:"is_public" description:"1- é funcionário público, 0- não é funcionário público"`
	SentWarning string  `json:"sent_warning" description:"Avisos enviados aos users"`
	//TODO: Isso pode se tornar o EMAIL da tabela warning
}
