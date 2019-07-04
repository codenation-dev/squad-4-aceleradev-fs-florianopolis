package model

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// Customer of the bank
type Customer struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" description:"Nome completo do cliente"`
	Wage        float32 `json:"wage" description:"Salário bruto mensal, sem os extras ocasionais (férias...)"`
	IsPublic    int8    `json:"is_public" description:"1- é funcionário público, 0- não é funcionário público"`
	SentWarning string  `json:"sent_warning" description:"Avisos enviados aos users"`
	//TODO: Isso pode se tornar o ID da tabela warning
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r)
}

// FormatString returns the string using the NeoWay's format
func FormatString(s string) string {
	b := make([]byte, len(s))
	fmt.Println(b)
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	_, _, err := t.Transform(b, []byte(s), true)
	if err != nil {
		panic(err)
	}
	b = bytes.Trim(b, "\x00") // Trim the null values
	sUpper := strings.ToUpper(string(b))

	return sUpper
}
