package fr_DZ

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "DZD", DisplayName: "", Symbol: "DA"},
	}
}
