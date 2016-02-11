package fr_BI

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "BIF", DisplayName: "", Symbol: "FBu"},
	}
}
