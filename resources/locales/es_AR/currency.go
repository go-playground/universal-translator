package es_AR

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "ARS", DisplayName: "", Symbol: "$"},
		{Currency: "GEL", DisplayName: "", Symbol: "GEL"},
		{Currency: "USD", DisplayName: "", Symbol: "US$"},
	}
}
