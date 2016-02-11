package en_SG

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "SGD", DisplayName: "", Symbol: "$"},
	}
}
