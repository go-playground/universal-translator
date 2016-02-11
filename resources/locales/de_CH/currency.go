package de_CH

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "BYR", DisplayName: "Weissrussischer Rubel", Symbol: ""},
		{Currency: "EUR", DisplayName: "", Symbol: "EUR"},
	}
}
