package bem

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "ZMW", DisplayName: "", Symbol: "K"},
	}
}
