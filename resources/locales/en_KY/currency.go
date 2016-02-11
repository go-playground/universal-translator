package en_KY

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "KYD", DisplayName: "", Symbol: "$"},
	}
}
