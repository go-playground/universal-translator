package fr_HT

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "HTG", DisplayName: "", Symbol: "G"},
	}
}
