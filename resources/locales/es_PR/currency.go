package es_PR

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "USD", DisplayName: "", Symbol: "$"},
	}
}
