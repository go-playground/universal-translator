package es_UY

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "USD", DisplayName: "", Symbol: "US$"},
		{Currency: "UYU", DisplayName: "", Symbol: "$"},
	}
}
