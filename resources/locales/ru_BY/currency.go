package ru_BY

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "BYR", DisplayName: "", Symbol: "р."},
		{Currency: "RUR", DisplayName: "", Symbol: "RUR"},
	}
}
