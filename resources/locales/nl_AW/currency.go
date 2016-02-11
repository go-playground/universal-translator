package nl_AW

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "AWG", DisplayName: "", Symbol: "Afl."},
	}
}
