package es_DO

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "DOP", DisplayName: "", Symbol: "RD$"},
		{Currency: "USD", DisplayName: "", Symbol: "US$"},
	}
}
