package ms_BN

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "BND", DisplayName: "", Symbol: "$"},
	}
}
