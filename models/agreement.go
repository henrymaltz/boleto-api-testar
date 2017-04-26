package models

import (
	"fmt"
	"regexp"

	"bitbucket.org/mundipagg/boletoapi/util"
)

// Agreement afiliação do cliente com o bano
type Agreement struct {
	AgreementNumber int
	Wallet          int16
	WalletVariation int16
	Agency          string
	AgencyDigit     string
	Account         string
	AccountDigit    string
}

// IsAgencyValid retorna se é uma agência válida
func (a *Agreement) IsAgencyValid() error {
	re := regexp.MustCompile("(\\D+)")
	ag := util.PadLeft(re.ReplaceAllString(a.Agency, ""), "0", 4)
	if len(ag) < 5 {
		a.Agency = ag
		return nil
	}
	return NewErrorResponse("MPAgency", "Agência inválida, deve conter até 4 dígitos")
}

// CalculateAgencyDigit calcula dígito da agência
func (a *Agreement) CalculateAgencyDigit(digitCalculator func(agency string) string) {
	re := regexp.MustCompile("(\\D+)")
	ad := re.ReplaceAllString(a.AgencyDigit, "")
	l := len(ad)
	if l < 2 && l > 0 {
		a.AgencyDigit = ad
		fmt.Println(a.AgencyDigit)
	} else {
		fmt.Println("hue")
		s := digitCalculator(a.AgencyDigit)
		a.AgencyDigit = s
		fmt.Println(a.AgencyDigit)
	}
}

// IsAccountValid retorna se é uma conta válida
func IsAccountValid(a *Agreement, accountLength uint) bool {
	re := regexp.MustCompile("(\\D+)")
	a.Account = re.ReplaceAllString(a.Account, "")
	a.Account = util.PadLeft(a.Account, "0", accountLength)
	return len(a.Account) < int(accountLength)
}

// IsAccountValid retorna se é uma conta válida
func (a Agreement) IsAccountValid(accountLength int) (string, error) {
	re := regexp.MustCompile("(\\D+)")
	ac := util.PadLeft(re.ReplaceAllString(a.Account, ""), "0", uint(accountLength))
	if len(ac) < accountLength+1 {
		return ac, nil
	}
	return "", NewErrorResponse("MPAccount", fmt.Sprintf("Conta inválida, deve conter até %d dígitos", accountLength))
}

// IsAccountDigitValid retorna se o dígito da conta é válido
func (a Agreement) IsAccountDigitValid() (string, error) {
	re := regexp.MustCompile("(\\D+)")
	ad := re.ReplaceAllString(a.AccountDigit, "")
	l := len(ad)
	if l < 2 && l > 0 {
		return ad, nil
	}
	return "", NewErrorResponse("MPAccountDigit", "Dígito da conta inválido. Deve conter apenas um dígito.")
}
