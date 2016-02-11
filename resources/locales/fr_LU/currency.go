package fr_LU

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "FRF", DisplayName: "", Symbol: "FRF"},
		{Currency: "LUF", DisplayName: "", Symbol: "F"},
	}
}
