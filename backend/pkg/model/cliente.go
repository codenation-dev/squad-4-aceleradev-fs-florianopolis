package model

type Cliente struct {
	ID           int     `json:"id_cliente"`
	Nome         string  `json:"nome"`
	NomePesquisa string  `json:"nome_pesquisa"`
	Salario      float64 `json:"salario"`
}
