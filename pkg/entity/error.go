package entity

import "errors"

var (
	// ErrEmptyFields is used when there is an empty field
	ErrEmptyFields = errors.New("um ou mais campos vazios")
	// ErrInvalidEmail is the error used if the mail is invalid
	ErrInvalidEmail    = errors.New("email inválido")
	ErrUserNotFound    = errors.New("usuário não encontrado")
	ErrDuplicatedUser  = errors.New("usuário já existe")
	ErrUnmarshal       = errors.New("problema ao decodificar os dados (Unmarshal)")
	ErrUnauthorized    = errors.New("erro na autenticação: usuário e/ou senha inválidos")
	ErrDownloadingFile = errors.New("erro ao fazer download do arquivo")

	ErrTableExists    = errors.New("tabela já existe")
	ErrNotImplemented = errors.New("função ainda não implementada")
)
