package ru_MD

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "MDL", DisplayName: "", Symbol: "L"},
	}
}
