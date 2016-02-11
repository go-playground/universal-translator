package en_SE

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "SEK", DisplayName: "", Symbol: "kr"},
	}
}
