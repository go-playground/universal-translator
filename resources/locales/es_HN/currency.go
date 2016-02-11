package es_HN

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "HNL", DisplayName: "", Symbol: "L"},
	}
}
