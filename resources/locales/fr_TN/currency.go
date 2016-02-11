package fr_TN

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "TND", DisplayName: "", Symbol: "DT"},
	}
}
