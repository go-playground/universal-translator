package ru_KG

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "KGS", DisplayName: "", Symbol: "сом"},
	}
}
