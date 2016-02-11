package fr_SY

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "SYP", DisplayName: "", Symbol: "LS"},
	}
}
