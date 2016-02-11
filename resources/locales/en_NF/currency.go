package en_NF

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "AUD", DisplayName: "", Symbol: "$"},
	}
}
