package fr_MU

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "MUR", DisplayName: "", Symbol: "Rs"},
	}
}
