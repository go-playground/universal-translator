package en_LR

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "LRD", DisplayName: "", Symbol: "$"},
	}
}
