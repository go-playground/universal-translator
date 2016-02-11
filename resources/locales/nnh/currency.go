package nnh

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "XAF", DisplayName: "feláŋ CFA", Symbol: "FCFA"},
	}
}
