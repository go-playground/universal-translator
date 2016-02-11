package es_US

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "JPY", DisplayName: "", Symbol: "¥"},
		{Currency: "USD", DisplayName: "", Symbol: "$"},
	}
}
