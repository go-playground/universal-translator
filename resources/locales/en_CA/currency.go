package en_CA

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "CAD", DisplayName: "", Symbol: "$"},
	}
}
