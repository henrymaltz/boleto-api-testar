package bank

import (
	"fmt"

	"github.com/mundipagg/boleto-api/bb"
	"github.com/mundipagg/boleto-api/bradescoNetEmpresa"
	"github.com/mundipagg/boleto-api/bradescoShopFacil"
	"github.com/mundipagg/boleto-api/caixa"
	"github.com/mundipagg/boleto-api/citibank"
	"github.com/mundipagg/boleto-api/itau"
	"github.com/mundipagg/boleto-api/log"
	"github.com/mundipagg/boleto-api/models"
	"github.com/mundipagg/boleto-api/santander"
)

//Bank é a interface que vai oferecer os serviços em comum entre os bancos
type Bank interface {
	ProcessBoleto(*models.BoletoRequest) (models.BoletoResponse, error)
	RegisterBoleto(*models.BoletoRequest) (models.BoletoResponse, error)
	ValidateBoleto(*models.BoletoRequest) models.Errors
	GetBankNumber() models.BankNumber
	Log() *log.Log
}

//Get retorna estrategia de acordo com o banco ou erro caso o banco não exista
func Get(number models.BankNumber) (Bank, error) {
	switch number {
	case models.BancoDoBrasil:
		return bb.New(), nil
	case models.BradescoShopFacil:
		return bradescoShopFacil.New(), nil
	case models.Caixa:
		return caixa.New(), nil
	case models.BradescoNetEmpresa:
		return bradescoNetEmpresa.New(), nil
	case models.Citibank:
		return citibank.New(), nil
	case models.Santander:
		return santander.New(), nil
	case models.Itau:
		return itau.New(), nil
	default:
		return nil, models.NewErrorResponse("MPBankNumber", fmt.Sprintf("Banco %d não existe", number))
	}
}
