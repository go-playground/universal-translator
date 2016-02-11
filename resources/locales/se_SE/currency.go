package se_SE

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "NOK", DisplayName: "", Symbol: "Nkr"},
		{Currency: "SEK", DisplayName: "", Symbol: "kr"},
	}
}
