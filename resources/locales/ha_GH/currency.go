package ha_GH

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "GHS", DisplayName: "", Symbol: "GH₵"},
	}
}
