package model

import "time"

// PublicFunc models the public employee profile with relevant informations
type Funcionario struct {
	ID              int        `json:"id_importaco,omitempty"`
	MesReferencia   *time.Time `json:"mes_referencia,omitempty"`
	Nome            string     `json:"nome"`
	NomePesquisa    string     `json:"nome_pesquisa,omitempty"`
	Cargo           string     `json:"cargo,omitempty"`
	Orgao           string     `json:"orgao,omitempty"`
	Estado          string     `json:"estado,omitempty"`
	SalarioMensal   float64    `json:"salario,omitempty"`
	SalarioFerias   float64    `json:"salario_ferias,omitempty"`
	PagtoEventual   float64    `json:"pagto_eventual,omitempty"`
	LicencaPremio   float64    `json:"licenca_premio,omitempty"`
	AbonoSalario    float64    `json:"abono_salario,omitempty"`
	RedutorSalarial float64    `json:"redutor_salarial,omitempty"`
	TotalLiquido    float64    `json:"total_liquido,omitempty"`
}
