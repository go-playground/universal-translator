package en_GG

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "GBP", DisplayName: "UK Pound", Symbol: ""},
	}
}
