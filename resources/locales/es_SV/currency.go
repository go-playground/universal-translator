package es_SV

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "USD", DisplayName: "", Symbol: "$"},
	}
}
