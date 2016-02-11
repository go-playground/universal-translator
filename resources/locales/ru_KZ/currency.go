package ru_KZ

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "KZT", DisplayName: "", Symbol: "₸"},
	}
}
