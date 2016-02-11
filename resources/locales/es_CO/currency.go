package es_CO

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "COP", DisplayName: "", Symbol: "$"},
		{Currency: "USD", DisplayName: "", Symbol: "US$"},
	}
}
