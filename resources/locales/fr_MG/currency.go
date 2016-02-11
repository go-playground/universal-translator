package fr_MG

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "MGA", DisplayName: "", Symbol: "Ar"},
	}
}
