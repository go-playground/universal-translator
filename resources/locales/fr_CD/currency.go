package fr_CD

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "CDF", DisplayName: "", Symbol: "FC"},
	}
}
