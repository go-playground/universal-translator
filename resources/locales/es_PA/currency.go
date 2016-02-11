package es_PA

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "PAB", DisplayName: "", Symbol: "B/."},
	}
}
