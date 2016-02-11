package fr_SC

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "SCR", DisplayName: "", Symbol: "SR"},
	}
}
