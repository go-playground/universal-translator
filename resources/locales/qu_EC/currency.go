package qu_EC

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "PEN", DisplayName: "", Symbol: "PEN"},
		{Currency: "USD", DisplayName: "", Symbol: "$"},
	}
}
