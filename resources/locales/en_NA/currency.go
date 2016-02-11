package en_NA

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "NAD", DisplayName: "", Symbol: "$"},
	}
}
