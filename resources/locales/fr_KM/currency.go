package fr_KM

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "KMF", DisplayName: "", Symbol: "CF"},
	}
}
