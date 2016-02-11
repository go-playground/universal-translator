package en_MT

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "GBP", DisplayName: "", Symbol: "GB£"},
	}
}
