package fr_RW

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "RWF", DisplayName: "", Symbol: "RF"},
	}
}
