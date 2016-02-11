package de_LU

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "LUF", DisplayName: "", Symbol: "F"},
	}
}
