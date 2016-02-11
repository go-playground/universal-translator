package kw

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "EUR", DisplayName: "Euro", Symbol: ""},
	}
}
