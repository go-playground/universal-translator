package es_CU

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "CUP", DisplayName: "", Symbol: "$"},
		{Currency: "USD", DisplayName: "", Symbol: "US$"},
	}
}
