package fr_VU

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "VUV", DisplayName: "", Symbol: "VT"},
	}
}
