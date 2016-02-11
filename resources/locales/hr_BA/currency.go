package hr_BA

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "BAM", DisplayName: "", Symbol: "KM"},
	}
}
