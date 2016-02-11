package en_WS

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "WST", DisplayName: "", Symbol: "WS$"},
	}
}
