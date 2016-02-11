package en_BW

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "BWP", DisplayName: "", Symbol: "P"},
	}
}
