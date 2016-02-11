package nl_SR

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "SRD", DisplayName: "", Symbol: "$"},
	}
}
