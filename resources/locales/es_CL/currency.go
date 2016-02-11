package es_CL

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "CLP", DisplayName: "Peso chileno", Symbol: "$"},
		{Currency: "USD", DisplayName: "", Symbol: "US$"},
	}
}
