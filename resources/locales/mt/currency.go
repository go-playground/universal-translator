package mt

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "EUR", DisplayName: "ewro", Symbol: "€"},
		{Currency: "MTL", DisplayName: "Lira Maltija", Symbol: ""},
		{Currency: "XXX", DisplayName: "Munita Mhux Magħrufa jew Mhux Valida", Symbol: ""},
	}
}
