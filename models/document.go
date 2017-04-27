package models

import (
	"regexp"
	"strings"
)

// Document nó com o tipo de documento e número do documento
type Document struct {
	Type   DocumentType
	Number DocumentNumber
}

// DocumentType o tipo de documento pode ser CPF ou CNPJ
type DocumentType string

// IsCpf diz se o DocumentType é um CPF
func (d DocumentType) IsCpf() bool {
	return strings.ToUpper(string(d)) == "CPF"
}

// IsCnpj diz se o DocumentType é um CNPJ
func (d DocumentType) IsCnpj() bool {
	return strings.ToUpper(string(d)) == "CNPJ"
}

// DocumentNumber o número do documento, poder ser um CPF ou CNPJ
type DocumentNumber string

// ValidateCPF verifica se é um CPF válido
func (d *Document) ValidateCPF() error {
	re := regexp.MustCompile("(\\D+)")
	cpf := re.ReplaceAllString(string(d.Number), "")
	if len(cpf) == 11 {
		d.Number = DocumentNumber(cpf)
		return nil
	}
	return ErrorResponse{Code: "MPDocumentNumber", Message: "CPF inválido"}
}

// IsCnpj verifica se é um Cnpj válido
func (d DocumentNumber) IsCnpj() bool {
	return len(d) == 14
}
